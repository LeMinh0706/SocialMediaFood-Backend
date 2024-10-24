package server

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/router"
	"github.com/LeMinh0706/SocialMediaFood-Backend/swag/docs"
	swaggerfiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(s *Server) {

	docs.SwaggerInfo.BasePath = "/api"
	a := s.Router.Group("/api")
	{
		router.NewUserRouter(s.Router, a, s.UserService)
	}

	s.Router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
