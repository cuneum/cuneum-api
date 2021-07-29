package model

import (
	"github.com/cuneum/cuneum-api/pkg/pb"
)

type ArticleList struct {
	Data       []Article
	Pagenation Pagenation
}

func (a *ArticleList) Into() *pb.ArticleList {
	result := &pb.ArticleList{
		Data:       []*pb.Article{},
		Pagenation: a.Pagenation.Into(),
	}
	for _, v := range a.Data {
		result.Data = append(result.Data, v.Into())
	}

	return result
}

type Article struct {
	ID          string
	Title       string
	URL         string
	ImageURL    string
	PublishedAt string
}

func (a *Article) Into() *pb.Article {
	return &pb.Article{
		Id:          a.ID,
		Title:       a.Title,
		Url:         a.URL,
		ImageUrl:    a.ImageURL,
		PublishedAt: a.PublishedAt,
	}
}
