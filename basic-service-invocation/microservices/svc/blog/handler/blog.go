package handler

import (
	"context"

	pb "github.com/kzmake/examples/basic-service-invocation/gen/go/blog/v1"
)

type blog struct {
	pb.UnimplementedBlogServiceServer
}

var _ pb.BlogServiceServer = new(blog)

func NewBlog() pb.BlogServiceServer {
	return &blog{}
}

func (h *blog) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	blog := &pb.Blog{
		Url: "https://example.com/" + req.GetUsername(),
	}

	// do something

	return &pb.CreateResponse{Blog: blog}, nil
}
