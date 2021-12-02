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
	CreateArticle(ctx context.Context, article *Article) error
}

type ArticleUsecase struct{
	repo ArticleRepo
}

func NewArticleUsecase(repo ArticleRepo) *ArticleUsecase {
	return &ArticleUsecase{
		repo: repo,
	}
}

func (uc *ArticleUsecase) List(ctx context.Context) (ps []*Article, err error) {
	ps, err = uc.repo.ListArticle(ctx)
	if err != nil {
		return
	}
	return
}

func (uc *ArticleUsecase) Create(ctx context.Context, article *Article) error {
	return uc.repo.CreateArticle(ctx, article)
}




