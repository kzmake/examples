package handler

import (
	"context"
	"fmt"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	adder "github.com/kzmake/examples/distributed-calculator/gen/go/adder/v1"
	pb "github.com/kzmake/examples/distributed-calculator/gen/go/subtractor/v1"
)

type subtractor struct {
	pb.UnimplementedSubtractorServer

	conn grpc.ClientConnInterface
}

var _ pb.SubtractorServer = new(subtractor)

func NewSubtractor(conn grpc.ClientConnInterface) pb.SubtractorServer { return &subtractor{conn: conn} }

func (h *subtractor) Sub(ctx context.Context, req *pb.SubRequest) (*pb.SubResponse, error) {
	operandTwo, err := strconv.ParseFloat(req.GetOperandTwo(), 64)
	if err != nil {
		return nil, err
	}

	c := adder.NewAdderClient(h.conn)
	res, err := c.Add(metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "svc-adder"), &adder.AddRequest{
		OperandOne: req.GetOperandOne(),
		OperandTwo: fmt.Sprintf("%f", -1.0*operandTwo),
	})
	if err != nil {
		return nil, err
	}

	return &pb.SubResponse{Result: res.GetResult()}, nil
}
