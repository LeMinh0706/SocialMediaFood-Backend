package post

import (
	"context"
	"fmt"
	"log"

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
	if description == "" && len(images) == 0 {
		return response.PostResponse{}, fmt.Errorf("Description or image can't empty")
	}

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

	postRes := response.PostRes(post, imgRes, user, post.DateCreatePost)

	return postRes, nil
}
