package main

import (
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/cmd/server"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	_ "github.com/lib/pq"
)

// @title SocialFood
// @version 2.0
// @description This is a sample server for Swagger.
// @host localhost:8070
// @BasePath /api/v1
func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}

	server, err := server.NewServer(config)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}
	server.Start(config.ServerAddress)
}
