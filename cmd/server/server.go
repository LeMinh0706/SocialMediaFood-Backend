package server

import (
	"fmt"

	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Server struct {
	Config      util.Config
	DBConn      *pgxpool.Pool
	TokenMaker  token.Maker
	RefeshMaker token.Maker
	Router      *gin.Engine
	Logger      *zap.Logger
}

func NewServer(db *pgxpool.Pool, config util.Config, logger *zap.Logger) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.SecretKey)
	if err != nil {
		return nil, fmt.Errorf("can not create token: %w", err)
	}
	refeshMaker, err := token.NewPasetoMaker(config.SecretKey)
	if err != nil {
		return nil, fmt.Errorf("can not create token: %w", err)
	}
	server := &Server{
		Config:      config,
		Router:      gin.Default(),
		TokenMaker:  tokenMaker,
		RefeshMaker: refeshMaker,
		DBConn:      db,
		Logger:      logger,
	}
	EnableCors(server.Router)
	server.Router.MaxMultipartMemory = 4 << 20
	server.NewRouter()
	return server, nil
}

func (server *Server) Start(address string) error {
	// server.Router.Use(middlewares.AuthorizeMiddleware(server.TokenMaker))
	return server.Router.Run(address)
}
