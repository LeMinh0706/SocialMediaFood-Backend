package post

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type PostResponse struct {
	ID           int64                `json:"id"`
	PostTypeID   int32                `json:"post_type_id"`
	AccountID    int64                `json:"account_id"`
	Description  string               `json:"description"`
	Lng          interface{}          `json:"lng"`
	Lat          interface{}          `json:"lat"`
	CreatedAt    time.Time            `json:"created_at"`
	Images       []db.PostImage       `json:"images"`
	Account      db.GetAccountByIdRow `json:"account"`
	ReactState   db.ReactPost         `json:"react_state"`
	TotalLike    int64                `json:"total_like"`
	TotalComment int64                `json:"total_comment"`
}

func PostRes(post db.CreatePostRow, account db.GetAccountByIdRow, imgs []db.PostImage, reactState db.ReactPost, totalLike, totalComment int64) PostResponse {
	return PostResponse{
		ID:           post.ID,
		PostTypeID:   post.PostTypeID,
		AccountID:    post.AccountID,
		Description:  post.Description.String,
		Lng:          post.Lng,
		Lat:          post.Lat,
		CreatedAt:    post.CreatedAt.Time,
		Images:       imgs,
		Account:      account,
		ReactState:   reactState,
		TotalLike:    totalLike,
		TotalComment: totalComment,
	}
}

func GetPostRes(post db.GetPostRow, account db.GetAccountByIdRow, imgs []db.PostImage, reactState db.ReactPost, totalLike, total_comment int64) PostResponse {
	return PostResponse{
		ID:           post.ID,
		PostTypeID:   post.PostTypeID,
		AccountID:    post.AccountID,
		Description:  post.Description.String,
		Lng:          post.Lng,
		Lat:          post.Lat,
		CreatedAt:    post.CreatedAt.Time,
		Images:       imgs,
		Account:      account,
		ReactState:   reactState,
		TotalLike:    totalLike,
		TotalComment: total_comment,
	}
}

// Util
func ConvertDescription(description string) pgtype.Text {
	if description == "" {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: description, Valid: true}
}

func CheckQuery(g *gin.Context, pageStr, pageSizeStr string) (int32, int32) {
	page, err := strconv.ParseInt(pageStr, 10, 32)
	if err != nil {
		response.ErrorResponse(g, 40001)
		return 0, 0
	}
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 32)
	if err != nil {
		response.ErrorResponse(g, 40002)
		return 0, 0
	}
	return int32(page), int32(pageSize)
}

func CheckValidPosition(g *gin.Context, lng, lat string) bool {
	if strings.TrimSpace(lng) == "" && strings.TrimSpace(lat) == "" {
		return true
	}
	_, err := strconv.ParseInt(lng, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 40020)
		return false
	}
	_, err = strconv.ParseInt(lat, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 40020)
		return false
	}
	if (lng == "" && lat != "") || (lng != "" && lat == "") {
		response.ErrorResponse(g, 40013)
		return false
	}
	return true
}

func AddImageFileError(g *gin.Context, number int, files []*multipart.FileHeader) ([]string, int) {
	const maxSize = 4 << 20
	if len(files) > number {
		return nil, 40005
	}
	for _, file := range files {
		if !util.FileExtCheck(file.Filename) {
			return nil, 40003
		}
	}

	var images []string
	for i, file := range files {
		if file.Size > maxSize {
			return nil, 41300
		}
		filename := fmt.Sprintf("upload/post/%d_%d%s", time.Now().Unix(), i, filepath.Ext(file.Filename))
		if err := g.SaveUploadedFile(file, filename); err != nil {
			return nil, 40021
		}
		images = append(images, filename)
	}
	return images, 201
}
