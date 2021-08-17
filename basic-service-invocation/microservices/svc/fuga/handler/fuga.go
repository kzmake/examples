package handler

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "github.com/kzmake/examples/basic-service-invocation/gen/go/fuga/v1"
	piyo "github.com/kzmake/examples/basic-service-invocation/gen/go/piyo/v1"
)

type fuga struct {
	pb.UnimplementedFugaServer

	conn grpc.ClientConnInterface
}

var _ pb.FugaServer = new(fuga)

func NewFuga(conn grpc.ClientConnInterface) pb.FugaServer {
	return &fuga{conn: conn}
}

func (h *fuga) Piyo(ctx context.Context, req *pb.PiyoRequest) (*pb.PiyoResponse, error) {
	c := piyo.NewPiyoClient(h.conn)
	res, err := c.Ping(metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "svc-piyo"), &piyo.PingRequest{})
	if err != nil {
		return nil, err
	}

	return &pb.PiyoResponse{Msg: strings.Join([]string{"called Fuga#Piyo", res.GetMsg()}, " : ")}, nil
}
