package handler

import (
	"context"
	crand "crypto/rand"
	"time"

	dapr "github.com/dapr/go-sdk/client"
	pb "github.com/kzmake/examples/basic-publish-and-subscribe/gen/go/user/v1"
	ulid "github.com/oklog/ulid/v2"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type user struct {
	pb.UnimplementedUserServiceServer
}

var _ pb.UserServiceServer = new(user)

func NewUser() pb.UserServiceServer { return &user{} }

func (h *user) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	created_at := time.Now()

	user := &pb.User{
		Id:   ulid.MustNew(ulid.Timestamp(created_at), crand.Reader).String(),
		Name: req.GetName(),
	}

	client, err := dapr.NewClient()
	if err != nil {
		return nil, err
	}

	data, err := proto.Marshal(&pb.CreatedEvent{
		User: &pb.CreatedEvent_User{
			Id:        user.GetId(),
			Name:      user.GetName(),
			CreatedAt: timestamppb.New(created_at),
		},
	})
	if err != nil {
		return nil, err
	}

	if err := client.PublishEvent(ctx, "pubsub", "user.v1.CreatedEvent", data); err != nil {
		return nil, err
	}

	return &pb.CreateResponse{User: user}, nil
}
