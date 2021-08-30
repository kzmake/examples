package handler

import (
	"context"
	"fmt"
	"strconv"

	pb "github.com/kzmake/examples/distributed-calculator/gen/go/divider/v1"
)

type divider struct {
	pb.UnimplementedDividerServer
}

var _ pb.DividerServer = new(divider)

func NewDivider() pb.DividerServer { return &divider{} }

func (h *divider) Div(ctx context.Context, req *pb.DivRequest) (*pb.DivResponse, error) {
	operandOne, err := strconv.ParseFloat(req.GetOperandOne(), 64)
	if err != nil {
		return nil, err
	}

	operandTwo, err := strconv.ParseFloat(req.GetOperandTwo(), 64)
	if err != nil {
		return nil, err
	}

	return &pb.DivResponse{Result: fmt.Sprintf("%f", operandOne/operandTwo)}, nil
}
