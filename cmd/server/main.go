package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/routers"
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
	fmt.Println(config.DBDriver)
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect database", err)
	}

	defer conn.Close()

	r := routers.NewRouter()
	r.Use(middlewares.CorsConfig())

	r.Run(":8070")

}
