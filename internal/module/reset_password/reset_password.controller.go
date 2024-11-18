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

func NewResetPasswordController(s IResetPasswordService, c util.Config, t token.Maker) ResetPasswordController {
	return ResetPasswordController{
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
// @Router       /reset-password [post]
func (rc *ResetPasswordController) ResetPasswordRequest(g *gin.Context) {
	email := g.Query("email")
	if strings.TrimSpace(email) == "" {
		response.ErrorResponse(g, 4000000)
		return
	}
	user, err := rc.service.ResetPasswordRequest(g, email)
	if err != nil {
		response.ErrorNonKnow(g, 401, err.Error())
		return
	}
	tokenId, _ := uuid.NewRandom()
	err = rc.service.AddRequestPassword(g, tokenId, user.ID, rc.config.AccessTokenDuration)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
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
