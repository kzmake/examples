package handler

import (
	"context"
	"fmt"
	"strconv"

	pb "github.com/kzmake/examples/distributed-calculator/gen/go/multiplier/v1"
)

type multiplier struct {
	pb.UnimplementedMultiplierServer
}

var _ pb.MultiplierServer = new(multiplier)

func NewMultiplier() pb.MultiplierServer { return &multiplier{} }

func (h *multiplier) Mul(ctx context.Context, req *pb.MulRequest) (*pb.MulResponse, error) {
	operandOne, err := strconv.ParseFloat(req.GetOperandOne(), 64)
	if err != nil {
		return nil, err
	}

	operandTwo, err := strconv.ParseFloat(req.GetOperandTwo(), 64)
	if err != nil {
		return nil, err
	}

	return &pb.MulResponse{Result: fmt.Sprintf("%f", operandOne*operandTwo)}, nil
}
