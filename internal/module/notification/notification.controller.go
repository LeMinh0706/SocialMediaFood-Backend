package notification

import (
	"strconv"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	service INotificationService
	token   token.Maker
}

func NewNotificationController(service INotificationService, token token.Maker) *NotificationController {
	return &NotificationController{
		service: service,
		token:   token,
	}
}

func (n *NotificationController) GetYourNotification(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	str := g.Param("id")
	account_id, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrAccountID)
		return
	}
	pageStr := g.Query("page")
	pageSizeStr := g.Query("page_size")
	page, pageSize := CheckQuery(g, pageStr, pageSizeStr)
	if page == 0 || pageSize == 0 {
		return
	}
	list, err := n.service.GetListNotification(g, auth.UserId, account_id, page, pageSize)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 200, list)
}
