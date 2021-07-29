package service

import (
	"context"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"

	"github.com/cuneum/cuneum-api/internal/model"
)

type ArticleService interface {
	GetArticleList(ctx context.Context) (*model.ArticleList, error)
}

type articleService struct {
}

func NewArticleService() ArticleService {
	return &articleService{}
}

func (a *articleService) GetArticleList(ctx context.Context) (*model.ArticleList, error) {
	fp := gofeed.NewParser()
	result, _ := fp.ParseURL(os.Getenv("RSS_URL"))

	data := []model.Article{}
	for _, v := range result.Items {
		data = append(data, model.Article{
			ID:          uuid.New().String(),
			Title:       v.Title,
			URL:         v.Link,
			ImageURL:    result.Items[0].Extensions["hatena"]["imageurl"][0].Value,
			PublishedAt: v.PublishedParsed.Format(time.RFC3339),
		})
	}

	return &model.ArticleList{
		Data:       data,
		Pagenation: model.Pagenation{},
	}, nil
}
