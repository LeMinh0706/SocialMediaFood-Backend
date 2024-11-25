package follower

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type AccountFollowResponse struct {
	Account db.GetAccountByIdRow `json:"account"`
	Status  string               `json:"status"`
}

type FollowResponse struct {
	FromFollow AccountFollowResponse `json:"from_follow"`
	ToFollow   AccountFollowResponse `json:"to_follow"`
}

type CreateFollowRequest struct {
	FromID int64 `json:"from_id"`
	ToID   int64 `json:"to_id"`
}

type ListFollow struct {
	Account []db.GetAccountByIdRow `json:"account"`
	Total   int64                  `json:"total"`
}
