package main

import (
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/cmd/server"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	_ "github.com/lib/pq"
)

// @title SocialFood Documentations
// @version 1.0
// @description This is SocialFood Swagger.
// @termsOfService github.com/LeMinh0706/SocialMediaFood-Backend

// @contact.name Đồ ăn công nghiệp (DACN)
// @contact.url github.com/LeMinh0706/SocialMediaFood-Backend
// @contact.email leminhken124356@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apiKey BearerAuth
// @in header
// @name Authorization
// @schema bearer

// @host localhost:8070
// @BasePath /api/v1
// @schema http
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
