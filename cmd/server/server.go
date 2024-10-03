package server

import (
	"fmt"

	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.SymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("Can not create token: %w", err)
	}
	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
	}

	server.NewRouter()
	return server, nil
}
