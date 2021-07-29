package handler

import (
	"context"

	"github.com/cuneum/cuneum-api/internal/service"
	"github.com/cuneum/cuneum-api/pkg/pb"
)

func NewArticle(articleService service.ArticleService) pb.ArticleServiceServer {
	return article{
		articleService: articleService,
	}
}

type article struct {
	pb.ArticleServiceServer

	articleService service.ArticleService
}

func (a article) List(ctx context.Context, _ *pb.PagenationRequest) (*pb.ArticleList, error) {
	result, err := a.articleService.GetArticleList(ctx)
	if err != nil {
		return nil, err
	}

	return result.Into(), nil
}
