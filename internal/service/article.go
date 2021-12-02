package service

import (
	"context"
	pb "myblog/api/back/v1"
	"github.com/go-kratos/kratos/v2/log"
	"myblog/internal/biz"
)

type BlogService struct {
	pb.UnimplementedBlogServiceServer

	log *log.Helper

	article *biz.ArticleUsecase
}

func NewBlogService(article *biz.ArticleUsecase) *BlogService {
	return &BlogService{
		article: article,
	}
}

func (s *BlogService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleReply, error)  {
	err := s.article.Create(ctx, &biz.Article{
		Title: "title",
		Content: "content",
	})
	return &pb.CreateArticleReply{}, err
}

func (s *BlogService) ListArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.ListArticleReply, error) {
	ps , err := s.article.list(ctx)
	reply := &pb.ListArticleReply{}
	for _, p := range ps {
		reply.Results = append(reply.Results, &pb.Article{
			Id:      p.ID,
			Title:   p.Title,
			Content: p.Content,
		})
	}
	return reply, err

}


