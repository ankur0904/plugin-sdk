package destination

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	pbBase "github.com/cloudquery/plugin-sdk/internal/pb/base/v0"
	pb "github.com/cloudquery/plugin-sdk/internal/pb/destination/v0"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedDestinationServer
	Plugin *destination.Plugin
	Logger zerolog.Logger
}

func (*Server) GetProtocolVersion(context.Context, *pbBase.GetProtocolVersion_Request) (*pbBase.GetProtocolVersion_Response, error) {
	return &pbBase.GetProtocolVersion_Response{
		Version: 2,
	}, nil
}

func (s *Server) Configure(ctx context.Context, req *pbBase.Configure_Request) (*pbBase.Configure_Response, error) {
	var spec specs.Destination
	if err := json.Unmarshal(req.Config, &spec); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to unmarshal spec: %v", err)
	}
	return &pbBase.Configure_Response{}, s.Plugin.Init(ctx, s.Logger, spec)
}

func (s *Server) GetName(context.Context, *pbBase.GetName_Request) (*pbBase.GetName_Response, error) {
	return &pbBase.GetName_Response{
		Name: s.Plugin.Name(),
	}, nil
}

func (s *Server) GetVersion(context.Context, *pbBase.GetVersion_Request) (*pbBase.GetVersion_Response, error) {
	return &pbBase.GetVersion_Response{
		Version: s.Plugin.Version(),
	}, nil
}

func (s *Server) Migrate(ctx context.Context, req *pb.Migrate_Request) (*pb.Migrate_Response, error) {
	var tables []*schema.Table
	if err := json.Unmarshal(req.Tables, &tables); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to unmarshal tables: %v", err)
	}
	return &pb.Migrate_Response{}, s.Plugin.Migrate(ctx, tables)
}

func (*Server) Write(pb.Destination_WriteServer) error {
	return status.Errorf(codes.Unimplemented, "method Write is deprecated please upgrade client")
}

// Note the order of operations in this method is important!
// Trying to insert into the `resources` channel before starting the reader goroutine will cause a deadlock.
func (s *Server) Write2(msg pb.Destination_Write2Server) error {
	resources := make(chan schema.DestinationResource)

	r, err := msg.Recv()
	if err != nil {
		if err == io.EOF {
			return msg.SendAndClose(&pb.Write2_Response{})
		}
		return status.Errorf(codes.Internal, "failed to receive msg: %v", err)
	}
	var tables schema.Tables
	if err := json.Unmarshal(r.Tables, &tables); err != nil {
		return status.Errorf(codes.InvalidArgument, "failed to unmarshal tables: %v", err)
	}
	var sourceSpec specs.Source
	if r.SourceSpec == nil {
		// this is for backward compatibility
		sourceSpec = specs.Source{
			Name: r.Source,
		}
	} else {
		if err := json.Unmarshal(r.SourceSpec, &sourceSpec); err != nil {
			return status.Errorf(codes.InvalidArgument, "failed to unmarshal source spec: %v", err)
		}
	}
	syncTime := r.Timestamp.AsTime()

	eg, ctx := errgroup.WithContext(msg.Context())
	eg.Go(func() error {
		return s.Plugin.Write(ctx, sourceSpec, tables, syncTime, resources)
	})

	for {
		r, err := msg.Recv()
		if err == io.EOF {
			close(resources)
			if err := eg.Wait(); err != nil {
				return status.Errorf(codes.Internal, "write failed: %v", err)
			}
			return msg.SendAndClose(&pb.Write2_Response{})
		}
		if err != nil {
			close(resources)
			if wgErr := eg.Wait(); wgErr != nil {
				return status.Errorf(codes.Internal, "failed to receive msg: %v and write failed: %v", err, wgErr)
			}
			return status.Errorf(codes.Internal, "failed to receive msg: %v", err)
		}
		var resource schema.DestinationResource
		if err := json.Unmarshal(r.Resource, &resource); err != nil {
			close(resources)
			if wgErr := eg.Wait(); wgErr != nil {
				return status.Errorf(codes.InvalidArgument, "failed to unmarshal resource: %v and write failed: %v", err, wgErr)
			}
			return status.Errorf(codes.InvalidArgument, "failed to unmarshal resource: %v", err)
		}
		select {
		case resources <- resource:
		case <-ctx.Done():
			close(resources)
			if err := eg.Wait(); err != nil {
				return status.Errorf(codes.Internal, "Context done: %v and failed to wait for plugin: %v", ctx.Err(), err)
			}
			return status.Errorf(codes.Internal, "Context done: %v", ctx.Err())
		}
	}
}

func (s *Server) GetMetrics(context.Context, *pb.GetDestinationMetrics_Request) (*pb.GetDestinationMetrics_Response, error) {
	stats := s.Plugin.Metrics()
	b, err := json.Marshal(stats)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal stats: %w", err)
	}
	return &pb.GetDestinationMetrics_Response{
		Metrics: b,
	}, nil
}

func (s *Server) DeleteStale(ctx context.Context, req *pb.DeleteStale_Request) (*pb.DeleteStale_Response, error) {
	var tables schema.Tables
	if err := json.Unmarshal(req.Tables, &tables); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to unmarshal tables: %v", err)
	}
	if err := s.Plugin.DeleteStale(ctx, tables, req.Source, req.Timestamp.AsTime()); err != nil {
		return nil, err
	}

	return &pb.DeleteStale_Response{}, nil
}

func (s *Server) Close(ctx context.Context, _ *pb.Close_Request) (*pb.Close_Response, error) {
	return &pb.Close_Response{}, s.Plugin.Close(ctx)
}