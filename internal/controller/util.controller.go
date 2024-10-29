package controller

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AddImageFileError(g *gin.Context, number int, files []*multipart.FileHeader) ([]string, error) {
	const maxSize = 4 << 20
	if len(files) > number {
		return nil, fmt.Errorf("images length shoud be less than 4")
	}
	for _, file := range files {
		if !util.FileExtCheck(file.Filename) {
			return nil, fmt.Errorf("only accept .jpeg/.jpg/.png/.gif")
		}
	}

	var images []string
	for i, file := range files {
		if file.Size > maxSize {
			return nil, fmt.Errorf("images size must less than 4 Mb")
		}
		filename := fmt.Sprintf("upload/post/%d_%d%s", time.Now().Unix(), i, filepath.Ext(file.Filename))
		if err := g.SaveUploadedFile(file, filename); err != nil {
			return nil, err
		}
		images = append(images, filename)
	}
	return images, nil
}

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
		response.ErrorNonKnow(g, 400, err.Error())
		return
	}
}
