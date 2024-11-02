package models

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

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

type AccountForPost struct {
	ID                   int64  `json:"id"`
	UserID               int64  `json:"user_id"`
	Fullname             string `json:"fullname"`
	UrlAvatar            string `json:"url_avatar"`
	UrlBackgroundProfile string `json:"url_background_profile"`
	RoleID               int32  `json:"role_id"`
}

func AccountPost(acc db.Account) AccountForPost {
	return AccountForPost{
		ID:                   acc.ID,
		UserID:               acc.UserID,
		Fullname:             acc.Fullname,
		UrlAvatar:            acc.UrlAvatar,
		UrlBackgroundProfile: acc.UrlBackgroundProfile,
		RoleID:               acc.RoleID,
	}
}

func AccountRes(account db.GetAccountByUserIdRow) AccountResponse {
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

func ListAccountResponse(all []db.GetAccountByUserIdRow) []AccountResponse {
	var list []AccountResponse
	for _, acc := range all {
		res := AccountRes(acc)
		list = append(list, res)
	}
	return list
}
