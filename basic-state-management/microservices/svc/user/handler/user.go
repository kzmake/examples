package handler

import (
	"context"
	crand "crypto/rand"

	pb "github.com/kzmake/examples/basic-state-management/gen/go/user/v1"
	ulid "github.com/oklog/ulid/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type user struct {
	pb.UnimplementedUserServiceServer

	userRepository UserRepository
}

var _ pb.UserServiceServer = new(user)

func NewUser() pb.UserServiceServer {
	return &user{userRepository: NewUserRepository()}
}

func (h *user) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	user := &User{
		ID:   UserID(ulid.MustNew(ulid.Now(), crand.Reader).String()),
		Name: UserName(req.GetName()),
	}

	if err := h.userRepository.Append(ctx, user); err != nil {
		return nil, err
	}

	u := &pb.User{
		Id:   string(user.ID),
		Name: req.GetName(),
	}

	// 201 Created
	header := metadata.Pairs("x-http-code", "201")
	_ = grpc.SendHeader(ctx, header)

	return &pb.CreateResponse{User: u}, nil
}

func (h *user) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	users, err := h.userRepository.List(ctx)
	if err != nil {
		return nil, err
	}

	us := []*pb.User{}
	for _, user := range users {
		u := &pb.User{
			Id:   string(user.ID),
			Name: string(user.Name),
		}
		us = append(us, u)
	}

	return &pb.ListResponse{Users: us}, nil
}

func (h *user) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	user, err := h.userRepository.Get(ctx, UserID(req.GetUserId()))
	if err != nil {
		return nil, err
	}

	u := &pb.User{
		Id:   string(user.ID),
		Name: string(user.Name),
	}

	return &pb.GetResponse{User: u}, nil
}

func (h *user) Rename(ctx context.Context, req *pb.RenameRequest) (*pb.RenameResponse, error) {
	user, err := h.userRepository.Get(ctx, UserID(req.GetUserId()))
	if err != nil {
		return nil, err
	}

	user.Name = UserName(req.GetName())

	if err := h.userRepository.Update(ctx, user); err != nil {
		return nil, err
	}

	updatedUser, err := h.userRepository.Get(ctx, UserID(req.GetUserId()))
	if err != nil {
		return nil, err
	}

	u := &pb.User{
		Id:   string(updatedUser.ID),
		Name: string(updatedUser.Name),
	}

	return &pb.RenameResponse{User: u}, nil
}

func (h *user) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	user, err := h.userRepository.Get(ctx, UserID(req.GetUserId()))
	if err != nil {
		return nil, err
	}

	if err := h.userRepository.Remove(ctx, user); err != nil {
		return nil, err
	}

	// 204 No Content
	header := metadata.Pairs("x-http-code", "204")
	_ = grpc.SendHeader(ctx, header)

	return &pb.DeleteResponse{}, nil
}
