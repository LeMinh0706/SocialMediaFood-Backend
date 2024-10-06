package post

import (
	"context"
	"fmt"
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
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

	var imgRes []db.PostImage

	for _, image := range images {
		i, err := p.postRepo.CreateImagePost(ctx, post.ID, image)
		if err != nil {
			return response.PostResponse{}, err
		}

		imgRes = append(imgRes, db.PostImage{
			ID:       i.ID,
			UrlImage: image,
			PostID:   post.ID,
		})

	}

	postRes := response.PostRes(post, imgRes, user, post.DateCreatePost)

	return postRes, nil
}

func (p *PostService) GetPost(ctx context.Context, post_id int64) (response.PostResponse, error) {

	post, err := p.postRepo.GetPost(ctx, post_id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return response.PostResponse{}, fmt.Errorf("NotFound")
		}
		return response.PostResponse{}, err
	}

	images, err := p.postRepo.GetImagePost(ctx, post_id)
	if err != nil {
		return response.PostResponse{}, err
	}

	user, err := user.NewUserService().GetUser(ctx, post.UserID)
	if err != nil {
		return response.PostResponse{}, err
	}

	res := response.PostRes(post, images, user, post.DateCreatePost)
	return res, nil

}

func (p *PostService) GetListPost(ctx context.Context, page, pageSize int64) ([]response.PostResponse, error) {

	posts, err := p.postRepo.GetListPost(ctx, int32(page), int32(pageSize))
	if err != nil {
		if err.Error() == "pq: OFFSET must not be negative" {
			return []response.PostResponse{}, fmt.Errorf("Cannot input zero here")
		}
		return []response.PostResponse{}, err
	}

	var res []response.PostResponse

	for _, post := range posts {
		images, err := p.postRepo.GetImagePost(ctx, post.ID)
		if err != nil {
			return []response.PostResponse{}, err
		}

		user, err := user.NewUserService().GetUser(ctx, post.UserID)
		if err != nil {
			return []response.PostResponse{}, err
		}
		posRes := response.PostRes(post, images, user, post.DateCreatePost)
		res = append(res, posRes)
	}
	if len(res) == 0 {
		return []response.PostResponse{}, nil //Đẹp trai ác
	}
	return res, nil
}
