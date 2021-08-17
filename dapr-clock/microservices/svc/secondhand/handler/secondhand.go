package handler

import (
	"context"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	pb "github.com/kzmake/examples/dapr-clock/gen/go/secondhand/v1"
)

type secondHand struct {
	pb.UnimplementedSecondHandServer

	conn *grpc.ClientConn
}

var _ pb.SecondHandServer = new(secondHand)

func NewSecondHand(conn *grpc.ClientConn) pb.SecondHandServer {
	return &secondHand{conn: conn}
}

func (h *secondHand) Now(ctx context.Context, req *pb.NowRequest) (*pb.NowResponse, error) {
	c := dapr.NewClientWithConnection(h.conn)

	item, err := c.GetState(ctx, "statestore-secondhand", "now")
	if err != nil {
		return nil, err
	}

	log.Debug().Msgf("item: %v", item)

	if item != nil && len(item.Value) != 0 {
		return &pb.NowResponse{Second: string(item.Value)}, nil
	}

	return &pb.NowResponse{Second: "-"}, nil
}
