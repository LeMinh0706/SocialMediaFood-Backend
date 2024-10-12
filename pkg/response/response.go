package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrSwaggerJson struct {
	Code   int         `json:"code"`
	Err    string      `json:"error"`
	Return interface{} `json:"return"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code, status int) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[status],
		Data:    nil,
	})
}

func ErrorNonKnow(c *gin.Context, code int, massage string) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: massage,
		Data:    nil,
	})
}
