package handler

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
)

func ValidateRegister(g *gin.Context, err error) {
	if validate, ok := err.(validator.ValidationErrors); ok {
		for _, vali := range validate {
			switch vali.Tag() {
			case "min":
				if vali.Field() == "Username" {
					response.ErrorResponse(g, 40008)
					return
				} else if vali.Field() == "Gender" {
					response.ErrorResponse(g, 40007)
					return
				} else if vali.Field() == "Password" {
					response.ErrorResponse(g, 40009)
					return
				} else if vali.Field() == "Fullname" {
					response.ErrorResponse(g, 40010)
					return
				}
			case "max":
				if vali.Field() == "Username" {
					response.ErrorResponse(g, 40008)
					return
				} else if vali.Field() == "Gender" {
					response.ErrorResponse(g, 40007)
					return
				}
			case "required":
				if vali.Field() == "Username" {
					response.ErrorResponse(g, 40008)
					return
				} else if vali.Field() == "Gender" {
					response.ErrorResponse(g, 40007)
					return
				} else if vali.Field() == "Password" {
					response.ErrorResponse(g, 40009)
					return
				} else if vali.Field() == "Fullname" {
					response.ErrorResponse(g, 40010)
					return
				}
			}
		}
	}
	response.ErrorResponse(g, 40000)
}

func SaveImage(g *gin.Context, type_image string, image *multipart.FileHeader) (string, int) {
	if !util.FileExtCheck(image.Filename) {
		return "", 40003
	}
	const maxSize = 4 << 20
	if image.Size > maxSize {
		return "", 41300
	}
	fileName := fmt.Sprintf("upload/%s/%d%s", type_image, time.Now().Unix(), filepath.Ext(image.Filename))
	if err := g.SaveUploadedFile(image, fileName); err != nil {
		return "", 40000
	}
	return fileName, 201
}

func CheckValidPosition(g *gin.Context, lng, lat string) bool {
	if lng == "" && lat == "" {
		return true
	}
	_, err := strconv.ParseFloat(lng, 64)
	if err != nil {
		response.ErrorResponse(g, 40020)
		return false
	}
	_, err = strconv.ParseFloat(lat, 64)
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

func FollowErr(g *gin.Context, err error) {
	switch err.Error() {
	case "not you":
		response.ErrorResponse(g, response.ErrYourSelf)
		return
	case "ERROR: duplicate key value violates unique constraint \"follower_to_follow_from_follow_idx\" (SQLSTATE 23505)":
		response.ErrorResponse(g, response.HaveFollow)
		return
	case "wait reply":
		response.ErrorResponse(g, response.AcceptForbidden)
		return
	case "no rows in result set":
		response.ErrorResponse(g, response.ErrAccountExists)
		return
	}
	response.ErrorNonKnow(g, 500, err.Error())
}

func ResetPasswordErr(g *gin.Context, err error) {
	switch err.Error() {
	case "not found":
		response.ErrorResponse(g, response.ErrEmailNotExists)
		return
	case "intime":
		response.ErrorResponse(g, response.YouHaveRequest)
		return
	case "request used":
		response.ErrorResponse(g, response.PasswordHaveChange)
		return
	}
	response.ErrorNonKnow(g, 500, err.Error())
}

func StatusCheck(g *gin.Context, status string) bool {
	a := []string{"accept", "friend", "request"}
	for _, s := range a {
		if status == s {
			return true
		}
	}
	response.ErrorResponse(g, 40415)
	return false
}

func CheckPostStringError(g *gin.Context, err error) {
	if err == pgx.ErrNoRows {
		response.ErrorResponse(g, response.ErrFindPost)
		return
	}
	if err.Error() == "not you" {
		response.ErrorResponse(g, response.ErrYourSelf)
		return
	}
	if err.Error() == "ERROR: duplicate key value violates unique constraint \"react_post_post_id_account_id_idx\" (SQLSTATE 23505)" {
		response.ErrorResponse(g, response.ErrLike)
		return
	}
	if err.Error() == "err like" {
		response.ErrorResponse(g, response.ErrUnlike)
		return
	}
	if err.Error() == "ERROR: duplicate key value violates unique constraint \"report_post_post_id_account_id_issue_id_idx\" (SQLSTATE 23505)" {
		response.ErrorResponse(g, response.ErrReport)
		return
	}
	response.ErrorNonKnow(g, 500, err.Error())
}
