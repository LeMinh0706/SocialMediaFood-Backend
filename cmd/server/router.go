package server

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/docs"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/hello"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/post"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/user"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (server *Server) NewRouter() {

	r := gin.Default()
	r.Static("upload", "./upload")
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		post.NewPostRouter(v1, server.tokenMaker)
		user.NewUserRouter(v1, server.tokenMaker)
		hello.NewHelloRouter(v1)
	}

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	server.router = r
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
