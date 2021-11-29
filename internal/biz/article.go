package biz

import (
	"context"
	"time"
)

type Article struct {
	ID        int64
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Like      int64
}

type ArticleRepo interface {
	ListArticle(ctx context.Context) ([]*Article, error)
	GetArticle(ctx context.Context, id int64) (*Article)
}

type ArticleUsecase struct{
	repo ArticleRepo
}

func NewArticleUsecase()  {

}



