package reset_password

import (
	"fmt"
	"strings"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ResetPasswordController struct {
	service IResetPasswordService
	config  util.Config
	token   token.Maker
}

func NewResetPasswordController(s IResetPasswordService, c util.Config, t token.Maker) *ResetPasswordController {
	return &ResetPasswordController{
		service: s,
		config:  c,
		token:   t,
	}
}

// ResetPassword godoc
// @Summary      Request link
// @Description  Mong chờ gì hơn ở chỗ này, chỉ cần nhập email thôi
// @Tags         ResetPassword
// @Accept       json
// @Produce      json
// @Param        email query string true "Your email"
// @Success      200  {object}  ResponseLink
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /forgot-password/request [post]
func (rc *ResetPasswordController) ForgotPassword(g *gin.Context) {
	email := g.Query("email")
	if strings.TrimSpace(email) == "" {
		response.ErrorResponse(g, 4000000)
		return
	}
	user, err := rc.service.ForgotPassword(g, email)
	if err != nil {
		ResetPasswordErr(g, err)
		return
	}
	tokenId, _ := uuid.NewRandom()
	err = rc.service.AddRequestPassword(g, tokenId, user.ID, rc.config.AccessTokenDuration)
	if err != nil {
		ResetPasswordErr(g, err)
		return
	}
	token, err := rc.token.CreateToken(tokenId, user.ID, user.Username, rc.config.AccessTokenDuration)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	link := fmt.Sprintf("%v?kamehameha=%v", rc.config.FrontEndUrl, token)
	time.Sleep(time.Second)
	response.SuccessResponse(g, 201, ResponseLink{Link: link})
}

func (rc *ResetPasswordController) FetchLink(g *gin.Context) {}

// ResetPassword godoc
// @Summary      Request link
// @Description  Notthing to say
// @Tags         ResetPassword
// @Accept       json
// @Produce      json
// @Param        request body ChangePasswordRequest true "request"
// @Success      201  {object}  ResponseLink
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /forgot-password/change [post]
func (rc *ResetPasswordController) ChangePassword(g *gin.Context) {
	var req ChangePasswordRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	payload, err := rc.token.VerifyToken(req.Token)
	if err != nil {
		response.ErrorResponse(g, response.ResetPasswordTimeOut)
		return
	}

	err = rc.service.ChangePassword(g, payload.Id, payload.UserId, req.NewPassword)
	if err != nil {
		ResetPasswordErr(g, err)
		return
	}

	response.SuccessResponse(g, response.ChangePassword, nil)
}