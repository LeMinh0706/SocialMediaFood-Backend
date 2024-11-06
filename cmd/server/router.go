package server

import (
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/factory"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/user"
	"github.com/LeMinh0706/SocialMediaFood-Backend/swag/docs"
	swaggerfiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) NewRouter() {
	fac, err := factory.NewFactory(s.DBConn)
	if err != nil {
		log.Fatal(err)
	}

	docs.SwaggerInfo.BasePath = "/api"
	a := s.Router.Group("/api")
	{
		Static(s.Router)
		user.NewUserRouter(s.Router, a, fac.UserService, s.TokenMaker, s.Config)
	}

	s.Router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
