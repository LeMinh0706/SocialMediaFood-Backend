package service

import (
	"context"
	"fmt"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
)

type PostService struct {
	postRepo    *repo.PostRepository
	userService *UserService
}

func NewPostService(repo *repo.PostRepository, userService *UserService) *PostService {

	return &PostService{
		postRepo:    repo,
		userService: userService,
	}
}

func (p *PostService) CreatePost(ctx context.Context, description string, user_id int64, images []string) (response.PostResponse, error) {
	var res response.PostResponse
	if description == "" && len(images) == 0 {
		return response.PostResponse{}, fmt.Errorf("Description or image can't empty")
	}

	if len(images) > 4 {
		return res, fmt.Errorf("Number of image can't more than 4")
	}

	post, err := p.postRepo.CreatePost(ctx, description, user_id)
	if err != nil {
		return res, err
	}

	user, err := p.userService.GetUser(ctx, user_id)
	if err != nil {
		return res, err
	}

	var imgRes []db.PostImage

	for _, image := range images {
		i, err := p.postRepo.CreateImagePost(ctx, post.ID, image)
		if err != nil {
			return res, err
		}

		imgRes = append(imgRes, db.PostImage{
			ID:       i.ID,
			UrlImage: image,
			PostID:   post.ID,
		})

	}

	res = response.PostRes(post, imgRes, user, post.DateCreatePost)

	return res, nil
}

func (p *PostService) GetPost(ctx context.Context, post_id int64) (response.PostResponse, error) {
	var res response.PostResponse
	post, err := p.postRepo.GetPost(ctx, post_id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return res, fmt.Errorf("NotFound")
		}
		return res, err
	}

	images, err := p.postRepo.GetImagePost(ctx, post_id)
	if err != nil {
		return res, err
	}

	user, err := p.userService.GetUser(ctx, post.UserID)
	if err != nil {
		return res, err
	}

	res = response.PostRes(post, images, user, post.DateCreatePost)
	return res, nil

}

func (p *PostService) GetListPost(ctx context.Context, page, pageSize int64) ([]response.PostResponse, error) {
	var res []response.PostResponse
	switch {
	case page <= 0:
		return res, fmt.Errorf("page need 1 or higher")
	case pageSize <= 0:
		return res, fmt.Errorf("page_size need 1 or higher")
	case pageSize > 10:
		return res, fmt.Errorf("page_size is lower than 10")
	}
	posts, err := p.postRepo.GetListPost(ctx, int32(page), int32(pageSize))
	if err != nil {
		return []response.PostResponse{}, err
	}

	for _, post := range posts {
		images, err := p.postRepo.GetImagePost(ctx, post.ID)
		if err != nil {
			return []response.PostResponse{}, err
		}

		user, err := p.userService.GetUser(ctx, post.UserID)
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

///////////////////////////////////////////
// 	// Test  goroutine
// func (p *PostService) CreatePost(ctx context.Context, description string, user_id int64, images []string) (response.PostResponse, error) {
// 	var res response.PostResponse
// 	if description == "" && len(images) == 0 {
// 		return response.PostResponse{}, fmt.Errorf("Description or image can't empty")
// 	}

// 	if len(images) > 4 {
// 		return res, fmt.Errorf("Number of image can't more than 4")
// 	}

// 	var wg sync.WaitGroup
// 	postChan := make(chan db.Post, 1)
// 	userChan := make(chan db.GetUserByIdRow, 1)
// 	errChan := make(chan error, 1)
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		post, err := p.postRepo.CreatePost(ctx, description, user_id)
// 		if err != nil {
// 			errChan <- err
// 			return
// 		}
// 		postChan <- post
// 	}()

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		user, err := p.userService.GetUser(ctx, user_id)
// 		if err != nil {
// 			errChan <- err
// 			return
// 		}
// 		userChan <- user
// 	}()

// 	go func() {
// 		wg.Wait()
// 		close(userChan)
// 		close(postChan)
// 		close(errChan)
// 	}()
// 	if err := <-errChan; err != nil {
// 		return res, err
// 	}
// 	post := <-postChan
// 	user := <-userChan

// 	var imgRes []db.PostImage

// 	for _, image := range images {
// 		i, err := p.postRepo.CreateImagePost(ctx, post.ID, image)
// 		if err != nil {
// 			return res, err
// 		}

// 		imgRes = append(imgRes, db.PostImage{
// 			ID:       i.ID,
// 			UrlImage: image,
// 			PostID:   post.ID,
// 		})

// 	}

// 	res = response.PostRes(post, imgRes, user, post.DateCreatePost)

// 	return res, nil
// }
