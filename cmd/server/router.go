package server

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/docs"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/routers"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (server *Server) NewRouter() {

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		routers.NewPostRouter(v1, server.tokenMaker)
		routers.NewUserRouter(v1, server.tokenMaker)
	}

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	server.router = r
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
