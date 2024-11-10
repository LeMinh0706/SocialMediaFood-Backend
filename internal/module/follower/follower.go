package follower

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/gin-gonic/gin"
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

func FollowErr(g *gin.Context, err error) {
	switch err.Error() {
	case "not you":
		response.ErrorResponse(g, 40103)
		return
	case "ERROR: duplicate key value violates unique constraint \"follower_to_follow_from_follow_idx\" (SQLSTATE 23505)":
		response.ErrorResponse(g, 40111)
		return
	case "no rows in result set":
		response.ErrorResponse(g, response.ErrAccountExists)
		return
	}
	response.ErrorNonKnow(g, 500, err.Error())
}
