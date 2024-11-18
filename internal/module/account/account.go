package account

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/gin-gonic/gin"
)

type EmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type AccountResponse struct {
	ID                   int64  `json:"id"`
	UserID               int64  `json:"user_id"`
	Fullname             string `json:"fullname"`
	UrlAvatar            string `json:"url_avatar"`
	UrlBackgroundProfile string `json:"url_background_profile"`
	Gender               int32  `json:"gender"`
	Country              string `json:"country"`
	Language             string `json:"language"`
	Address              string `json:"address"`
	IsUpgrade            bool   `json:"is_upgrade"`
}

type UpdateNameReq struct {
	Fullname string `json:"fullname"`
}

func AccountRes(account db.Account) AccountResponse {
	return AccountResponse{
		ID:                   account.ID,
		UserID:               account.UserID,
		Fullname:             account.Fullname,
		UrlAvatar:            account.UrlAvatar,
		UrlBackgroundProfile: account.UrlBackgroundProfile,
		Gender:               account.Gender.Int32,
		Country:              account.Country.String,
		Language:             account.Language.String,
		Address:              account.Address.String,
		IsUpgrade:            account.IsUpgrade.Bool,
	}
}

func ListAccountResponse(all []db.Account) []AccountResponse {
	var list []AccountResponse
	for _, acc := range all {
		res := AccountRes(acc)
		list = append(list, res)
	}
	return list
}

func SaveAccountImage(g *gin.Context, type_image string, image *multipart.FileHeader) (string, int) {
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
