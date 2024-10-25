package server

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/factory"
)

func (server *Server) InitService() error {
	factory, err := factory.NewFactory(server.DBConn)
	if err != nil {
		return err
	}
	server.UserService = factory.UserService
	server.AccountService = factory.AccountService
	return nil
}
