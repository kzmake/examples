package handler

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "github.com/kzmake/examples/daprclock/gen/go/clock/v1"
	hourhand "github.com/kzmake/examples/daprclock/gen/go/hourhand/v1"
	secondhand "github.com/kzmake/examples/daprclock/gen/go/secondhand/v1"
)

type clock struct {
	pb.UnimplementedClockServer

	conn grpc.ClientConnInterface
}

var _ pb.ClockServer = new(clock)

func NewClock(conn grpc.ClientConnInterface) pb.ClockServer {
	return &clock{conn: conn}
}

func (h *clock) Now(ctx context.Context, req *pb.NowRequest) (*pb.NowResponse, error) {
	hourHandC := hourhand.NewHourHandClient(h.conn)
	hourHandR, err := hourHandC.Now(metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "svc-hourhand"), &hourhand.NowRequest{})
	if err != nil {
		return nil, err
	}

	secondHandC := secondhand.NewSecondHandClient(h.conn)
	secondHandR, err := secondHandC.Now(metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "svc-secondhand"), &secondhand.NowRequest{})
	if err != nil {
		return nil, err
	}

	return &pb.NowResponse{Hour: hourHandR.GetHour(), Minute: "34", Second: secondHandR.GetSecond()}, nil
}
