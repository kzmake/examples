package handler

import (
	"context"

	"github.com/dapr/go-sdk/service/common"
	pb "github.com/kzmake/examples/basic-publish-and-subscribe/gen/go/user/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

type wallet struct{}

func NewWallet() *wallet { return &wallet{} }

func (h *wallet) Create(ctx context.Context, in *common.TopicEvent) (bool, error) {
	log.Debug().Msgf("handle Wallet.Create")

	e := &pb.CreatedEvent{}
	if err := proto.Unmarshal(in.Data.([]byte), e); err != nil {
		return false, err
	}

	log.Debug().Msgf("recv event: %v", e)
	// do something

	return false, nil
}
