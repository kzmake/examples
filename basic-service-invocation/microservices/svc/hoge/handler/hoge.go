package handler

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	fuga "github.com/kzmake/examples/basic-service-invocation/gen/go/fuga/v1"
	pb "github.com/kzmake/examples/basic-service-invocation/gen/go/hoge/v1"
)

type hoge struct {
	pb.UnimplementedHogeServer

	conn grpc.ClientConnInterface
}

var _ pb.HogeServer = new(hoge)

func NewHoge(conn grpc.ClientConnInterface) pb.HogeServer {
	return &hoge{conn: conn}
}

func (h *hoge) Fuga(ctx context.Context, req *pb.FugaRequest) (*pb.FugaResponse, error) {
	c := fuga.NewFugaClient(h.conn)
	res, err := c.Piyo(metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "svc-fuga"), &fuga.PiyoRequest{})
	if err != nil {
		return nil, err
	}

	return &pb.FugaResponse{Msg: strings.Join([]string{"called Fuga#Piyo", res.GetMsg()}, " : ")}, nil
}
