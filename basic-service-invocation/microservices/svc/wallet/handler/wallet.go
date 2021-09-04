package handler

import (
	"context"
	crand "crypto/rand"
	"time"

	pb "github.com/kzmake/examples/basic-service-invocation/gen/go/wallet/v1"
	"github.com/oklog/ulid/v2"
)

type wallet struct {
	pb.UnimplementedWalletServiceServer
}

var _ pb.WalletServiceServer = new(wallet)

func NewWallet() pb.WalletServiceServer { return &wallet{} }

func (h *wallet) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	wallet := &pb.Wallet{
		Id: ulid.MustNew(ulid.Timestamp(time.Now()), crand.Reader).String(),
	}

	// do something

	return &pb.CreateResponse{Wallet: wallet}, nil
}
