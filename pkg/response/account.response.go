package response

import "github.com/LeMinh0706/SocialMediaFood-Backend/db"

type AccountResponse struct {
	ID                   int64       `json:"id"`
	UserID               int64       `json:"user_id"`
	Fullname             string      `json:"fullname"`
	UrlAvatar            string      `json:"url_avatar"`
	UrlBackgroundProfile string      `json:"url_background_profile"`
	Gender               int32       `json:"gender"`
	Country              string      `json:"country"`
	Language             string      `json:"language"`
	Address              string      `json:"address"`
	IsUpgrade            bool        `json:"is_upgrade"`
	Lng                  interface{} `json:"lng"`
	Lat                  interface{} `json:"lat"`
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
		Lng:                  account.Lng,
		Lat:                  account.Lat,
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
