package hello

import "github.com/gin-gonic/gin"

func NewHelloRouter(router *gin.RouterGroup) {
	userGroup := router.Group("/hello")
	uc := NewHelloController()
	{
		userGroup.GET(":name", uc.GetHelloParam)
		userGroup.GET("", uc.GetHelloQuery)
	}
}
