package handler

import (
	"context"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Tick interface {
	Tick(context.Context, *common.BindingEvent) ([]byte, error)
}

type tick struct {
	conn *grpc.ClientConn
}

func NewTick(conn *grpc.ClientConn) Tick {
	return &tick{conn: conn}
}

func (h *tick) Tick(ctx context.Context, in *common.BindingEvent) ([]byte, error) {
	log.Debug().Msgf("binding-1s: %v", in.Metadata)

	c := dapr.NewClientWithConnection(h.conn)
	if err := c.PublishEvent(ctx, "pubsub-x", "Ticked1s", nil); err != nil {
		return nil, err
	}

	return nil, nil
}
