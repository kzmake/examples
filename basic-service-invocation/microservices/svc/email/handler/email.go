package handler

import (
	"context"

	pb "github.com/kzmake/examples/basic-service-invocation/gen/go/email/v1"
)

type email struct {
	pb.UnimplementedEmailServiceServer
}

var _ pb.EmailServiceServer = new(email)

func NewEmail() pb.EmailServiceServer {
	return &email{}
}

func (h *email) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	email := &pb.Email{
		Address: req.GetUsername() + "@example.com",
	}

	// do something

	return &pb.CreateResponse{Email: email}, nil
}
