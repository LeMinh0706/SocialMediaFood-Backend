package gapi

import (
	"fmt"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/pb"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
)

type Server struct {
	pb.UnimplementedSocialMediaFoodServer
	Config         util.Config
	TokenMaker     token.Maker
	UserService    *service.UserService
	PostService    *service.PostService
	CommentService *service.CommentService
}

func NewServer(config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.SymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("Can not create token: %w", err)
	}
	server := &Server{
		Config:     config,
		TokenMaker: tokenMaker,
	}
	if err != nil {
		return nil, err
	}

	err = server.InitService()
	if err != nil {
		return nil, err
	}

	return server, nil
}
