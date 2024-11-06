package models

import "github.com/LeMinh0706/SocialMediaFood-Backend/db"

type FollowRequest struct {
	FromFollow int64 `json:"from_follow"`
	ToFollow   int64 `json:"to_follow"`
}

type FollowRespone struct {
	FromFollow db.Follower `json:"from_follow"`
	ToFollow   db.Follower `json:"to_follow"`
}
