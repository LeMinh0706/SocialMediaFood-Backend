package main

import (
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/cmd/server"
	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	_ "github.com/lib/pq"
)

// @title Foodioo Documentations
// @version 1.0
// @description This is SocialFood Swagger.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apiKey BearerAuth
// @in header
// @name Authorization
// @schema bearer

// @host localhost:8070
// @BasePath /api
// @schema http
func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}

	pg, err := db.GetDBConnection(config)
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()

	server, err := server.NewServer(pg, config)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}
	server.Start(config.ServerAddress)
}
