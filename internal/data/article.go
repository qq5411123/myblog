package data

import (
	"context"
	"myblog/internal/biz"
)

type articleRepo struct {
	data *Data
}

func NewArticleRepo(data *Data) biz.ArticleRepo {
	return &articleRepo{
		data : data,
	}
}

func (ar *articleRepo) ListArticle(ctx context.Context) ([]*biz.Article, error) {
	ps, err := ar.data.db.Article.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Article, 0)
	for _, p := range ps {
		rv = append(rv, &biz.Article{
			ID:        p.ID,
			Title:     p.Title,
			Content:   p.Content,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}
	return rv, nil
}

func (ar *articleRepo) CreateArticle(ctx context.Context, article *biz.Article) error {
	_, err := ar.data.db.Article.
		Create().
		SetTitle(article.Title).
		SetContent(article.Content).
		Save(ctx)
	return err
}
