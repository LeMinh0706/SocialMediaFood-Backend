package main

import (
	"net/http"

	docs "github.com/LeMinh0706/SocialMediaFood-Backend/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @Router /example/ping [get]
func Pong(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"status":  "200",
	})
}

// HelloExample godoc
// @Summary ping example
// @Schemes
// @Description just hello
// @Tags example
// @Accept json
// @Produce json
// @Router /example/hello [get]
func Hello(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "Hello",
		"status":  "200",
	})
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/ping", Pong)
			eg.GET("/hello", Hello)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8070")

}
