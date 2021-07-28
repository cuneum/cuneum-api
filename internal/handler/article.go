package handler

import (
	"context"
	"time"

	"github.com/cuneum/cuneum-api/pkg/pb"
	"github.com/google/uuid"
)

func NewArticle() pb.ArticleServiceServer {
	return article{}
}

type article struct {
	pb.ArticleServiceServer
}

func (a article) List(context.Context, *pb.PagenationRequest) (*pb.ArticleList, error) {
	return &pb.ArticleList{
		Data: []*pb.Article{
			{
				Id:          uuid.New().String(),
				Title:       "hoge",
				Url:         "https://example.com/hoge",
				ImageUrl:    "https://example.com/hoge.png",
				PublishedAt: time.Now().Format(time.RFC3339),
			},
		},
		Pagenation: &pb.PagenationResponse{},
	}, nil
}
