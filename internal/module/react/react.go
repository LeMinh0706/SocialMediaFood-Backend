package react

import "github.com/LeMinh0706/SocialMediaFood-Backend/db"

type ReactRequest struct {
	PostID    int64 `json:"post_id"`
	AccountID int64 `json:"account_id"`
}

type ReactResponse struct {
	PostID  int64                `json:"post_id"`
	State   int32                `json:"state"`
	Account db.GetAccountByIdRow `json:"account"`
}

type ListReactResponse struct {
	React []ReactResponse `json:"react"`
	Total int64           `json:"total"`
}
