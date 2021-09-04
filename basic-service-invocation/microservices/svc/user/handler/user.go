package handler

import (
	"context"
	crand "crypto/rand"
	"time"

	blogpb "github.com/kzmake/examples/basic-service-invocation/gen/go/blog/v1"
	emailpb "github.com/kzmake/examples/basic-service-invocation/gen/go/email/v1"
	pb "github.com/kzmake/examples/basic-service-invocation/gen/go/user/v1"
	walletpb "github.com/kzmake/examples/basic-service-invocation/gen/go/wallet/v1"
	ulid "github.com/oklog/ulid/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type user struct {
	pb.UnimplementedUserServiceServer

	conn grpc.ClientConnInterface
}

var _ pb.UserServiceServer = new(user)

func NewUser(conn grpc.ClientConnInterface) pb.UserServiceServer {
	return &user{conn: conn}
}

func (h *user) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	blog := blogpb.NewBlogServiceClient(h.conn)
	blogRes, err := blog.Create(
		metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "svc-blog"),
		&blogpb.CreateRequest{
			Username: req.GetName(),
		},
	)
	if err != nil {
		return nil, err
	}

	email := emailpb.NewEmailServiceClient(h.conn)
	emailRes, err := email.Create(
		metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "svc-email"),
		&emailpb.CreateRequest{
			Username: req.GetName(),
		},
	)
	if err != nil {
		return nil, err
	}

	wallet := walletpb.NewWalletServiceClient(h.conn)
	walletRes, err := wallet.Create(
		metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "svc-wallet"),
		&walletpb.CreateRequest{},
	)
	if err != nil {
		return nil, err
	}

	user := &pb.User{
		Id:           ulid.MustNew(ulid.Timestamp(time.Now()), crand.Reader).String(),
		Name:         req.GetName(),
		BlogUrl:      blogRes.GetBlog().GetUrl(),
		EmailAddress: emailRes.GetEmail().GetAddress(),
		WalletId:     walletRes.GetWallet().GetId(),
	}
	return &pb.CreateResponse{User: user}, nil
}
