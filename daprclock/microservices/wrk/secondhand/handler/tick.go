package handler

import (
	"context"
	"strconv"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Tick interface {
	Tick(context.Context, *common.TopicEvent) (bool, error)
}

type tick struct {
	conn *grpc.ClientConn
}

func NewTick(conn *grpc.ClientConn) Tick {
	return &tick{conn: conn}
}

func (h *tick) Tick(ctx context.Context, in *common.TopicEvent) (bool, error) {
	log.Debug().Msgf("pubsub-x#Ticked1s: %v", in)

	c := dapr.NewClientWithConnection(h.conn)

	item, err := c.GetState(ctx, "statestore-secondhand", "now")
	if err != nil {
		return false, err
	}

	log.Debug().Msgf("item: %v", item)
	var now int
	if item != nil && len(item.Value) != 0 {
		now, err = strconv.Atoi(string(item.Value))
		if err != nil {
			return false, err
		}
	} else {
		log.Info().Msgf("init")
		now = 0
	}

	now = now + 1
	if now >= 60 {
		now = 0
	}

	if err := c.SaveState(ctx, "statestore-secondhand", "now", []byte(strconv.Itoa(now))); err != nil {
		return false, err
	}

	if now == 0 {
		if err := c.PublishEvent(ctx, "pubsub-x", "Ticked60s", nil); err != nil {
			return false, err
		}
	}

	return false, nil
}
