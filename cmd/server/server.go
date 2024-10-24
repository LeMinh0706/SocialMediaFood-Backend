package server

import (
	"fmt"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Config      util.Config
	TokenMaker  token.Maker
	Router      *gin.Engine
	UserService *service.UserService
}

func NewServer(config util.Config) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.SecretKey)
	if err != nil {
		return nil, fmt.Errorf("can not create token: %w", err)
	}
	server := &Server{
		Config:     config,
		Router:     gin.Default(),
		TokenMaker: tokenMaker,
	}
	err = server.InitService()
	if err != nil {
		return nil, err
	}
	Static(server.Router)
	NewRouter(server)
	return server, nil
}

func (server *Server) Start(address string) error {
	// server.Router.Use(middlewares.AuthorizeMiddleware(server.TokenMaker))
	return server.Router.Run(address)
}
