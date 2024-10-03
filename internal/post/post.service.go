package post

import (
	"context"
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type PostService struct {
	postRepo *PostRepository
}

func NewPostService() *PostService {
	repo, err := NewPostRepo()
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
