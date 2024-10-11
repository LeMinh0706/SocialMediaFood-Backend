package server

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/factory"
)

func (server *Server) InitService() error {
	factory, err := factory.NewFactory()
	if err != nil {
		return err
	}
	server.UserService = factory.UserService
	server.PostService = factory.PostService
	server.CommentService = factory.CommentService
	server.ReactService = factory.ReactService
	return nil
}
