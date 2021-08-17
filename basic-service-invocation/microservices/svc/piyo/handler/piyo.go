package handler

import (
	"context"

	pb "github.com/kzmake/examples/basic-service-invocation/gen/go/piyo/v1"
)

type piyo struct {
	pb.UnimplementedPiyoServer
}

var _ pb.PiyoServer = new(piyo)

func NewPiyo() pb.PiyoServer { return &piyo{} }

func (h *piyo) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Msg: "called Piyo#Ping"}, nil
}
