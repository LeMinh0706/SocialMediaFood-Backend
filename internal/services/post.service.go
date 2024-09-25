package services

import (
	"context"
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
)

type PostService struct {
	postRepo *repo.PostRepository
}

func NewPostService() *PostService {
	repo, err := repo.NewPostRepo()
	if err != nil {
		log.Fatal("Error:", err)
	}
	return &PostService{
		postRepo: repo,
	}
}

func (p *PostService) CreatePost(ctx context.Context, description string, user_id int64) (db.Post, error) {
	post, err := p.postRepo.CreatePost(ctx, description, user_id)
	if err != nil {
		return db.Post{}, err
	}
	return post, nil
}
