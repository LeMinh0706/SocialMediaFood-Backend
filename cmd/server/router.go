package server

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/docs"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/router"
	swaggerfiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(s *Server) {

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := s.Router.Group("/api/v1")
	{
		router.NewUserRouter(s.Router, v1, s.TokenMaker, s.UserService, s.Config)
		router.NewPostRouter(s.Router, v1, s.TokenMaker, s.PostService)
		router.NewCommentRouter(s.Router, v1, s.TokenMaker, s.CommentService)
		router.NewReactRouter(s.Router, v1, s.TokenMaker, s.ReactService)
	}

	s.Router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
