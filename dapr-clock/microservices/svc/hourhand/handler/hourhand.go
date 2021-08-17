package handler

import (
	"context"

	pb "github.com/kzmake/examples/dapr-clock/gen/go/hourhand/v1"
)

type hourHand struct {
	pb.UnimplementedHourHandServer
}

var _ pb.HourHandServer = new(hourHand)

func NewHourHand() pb.HourHandServer { return &hourHand{} }

func (h *hourHand) Now(ctx context.Context, req *pb.NowRequest) (*pb.NowResponse, error) {
	return &pb.NowResponse{Hour: "12"}, nil
}
