package handler

import (
	"context"
	"fmt"
	"strconv"

	pb "github.com/kzmake/examples/distributed-calculator/gen/go/adder/v1"
)

type adder struct {
	pb.UnimplementedAdderServer
}

var _ pb.AdderServer = new(adder)

func NewAdder() pb.AdderServer { return &adder{} }

func (h *adder) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	operandOne, err := strconv.ParseFloat(req.GetOperandOne(), 64)
	if err != nil {
		return nil, err
	}

	operandTwo, err := strconv.ParseFloat(req.GetOperandTwo(), 64)
	if err != nil {
		return nil, err
	}

	return &pb.AddResponse{Result: fmt.Sprintf("%f", operandOne+operandTwo)}, nil
}
