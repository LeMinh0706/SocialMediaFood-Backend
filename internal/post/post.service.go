package post

import (
	"context"
	"log"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/user"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
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

func (p *PostService) CreatePost(ctx context.Context, description string, user_id int64, images []string) (response.PostResponse, error) {
	post, err := p.postRepo.CreatePost(ctx, description, user_id)
	if err != nil {
		return response.PostResponse{}, err
	}

	user, err := user.NewUserService().GetUser(ctx, user_id)
	if err != nil {
		return response.PostResponse{}, err
	}

	var imgRes []response.ImageResponse

	for _, image := range images {
		i, err := p.postRepo.CreateImagePost(ctx, post.ID, image)
		if err != nil {
			return response.PostResponse{}, err
		}

		imgRes = append(imgRes, response.ImageResponse{
			ID:       i.ID,
			UrlImage: image,
			PostId:   post.ID,
		})

	}

	postRes := response.PostResponse{
		ID:             post.ID,
		PostTypeID:     post.PostTypeID,
		UserID:         user_id,
		Description:    description,
		Images:         imgRes,
		User:           user,
		DateCreatePost: time.Now().Unix(),
	}

	return postRes, nil
}
