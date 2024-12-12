package notification

import (
	"strconv"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/handler"
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

// Notification godoc
// @Summary      Get list Notification
// @Description  Get list Notification with your account_id, page and page size (Limit-Offset)
// @Tags         Notification
// @Accept       json
// @Produce      json
// @Param        id path int true "AccountID"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Security BearerAuth
// @Success      200  {object}  []NotificationResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /notification/{id} [get]
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
	page, pageSize := handler.CheckQuery(g, pageStr, pageSizeStr)
	if page == 0 || pageSize == 0 {
		return
	}
	list, err := n.service.GetListNotification(g, auth.Username, account_id, page, pageSize)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 200, list)
}

// Notification godoc
// @Summary      Update Notification
// @Description  Update notification id
// @Tags         Notification
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Security BearerAuth
// @Success      201  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /notification/{id} [put]
func (n *NotificationController) SeenYourNotification(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	idStr := g.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrBadRequestId)
		return
	}
	err = n.service.IsSeen(g, auth.Username, id)
	if err != nil {
		handler.CheckPostStringError(g, err)
		return
	}
	response.SuccessResponse(g, 201, nil)
}

// Notification godoc
// @Summary      Update Notification
// @Description  Update seen list notification with your account_id
// @Tags         Notification
// @Accept       json
// @Produce      json
// @Param        id path int true "AccountID"
// @Security BearerAuth
// @Success      201  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /notification/seen-all/{id} [put]
func (n *NotificationController) SeenAllNoti(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	accStr := g.Param("id")
	id, err := strconv.ParseInt(accStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrBadRequestId)
		return
	}
	err = n.service.IsSeenAll(g, auth.Username, id)
	if err != nil {
		// handler.CheckPostStringError(g, err)
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 201, nil)
}

// Notification godoc
// @Summary      Delete Notification
// @Description  Soft delete notification and handle when you get list
// @Tags         Notification
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /notification/{id} [post]
func (n *NotificationController) DeleteNotification(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	idStr := g.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrBadRequestId)
		return
	}
	err = n.service.DeleteNoti(g, auth.Username, id)
	if err != nil {
		handler.CheckPostStringError(g, err)
		return
	}
	response.SuccessResponse(g, response.DeleteNoti, nil)
}
