package routers

import (
	"net/http"

	"github.com/LeMinh0706/SocialMediaFood-Backend/docs"
	c "github.com/LeMinh0706/SocialMediaFood-Backend/internal/controllers"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

//@title Social Media Food
// @BasePath /api/v1

// PingExample godoc
// @Summary ping pop
// @Schemes
// @Description do pong
// @Tags pong
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @Security Bearer
// @Router /pop/ping [get]
func Pong(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"status":  "200",
	})
}

func NewRouter() *gin.Engine {

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/hello/:name", c.NewHelloController().GetHelloParam)
			eg.GET("/hello", c.NewHelloController().GetHelloQuery)
			eg.POST("/hello", c.NewHelloController().PostHelloBody)
		}

		ag := v1.Group("/pop")
		{
			ag.GET("/ping", Pong)
		}
	}
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
