package main

import (
	"log"
	"net"

	"github.com/LeMinh0706/SocialMediaFood-Backend/cmd/gapi"
	"github.com/LeMinh0706/SocialMediaFood-Backend/cmd/server"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/pb"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}
	// GinRun(config)
	GrpcRun(config)
}

// @title SocialFood
// @version 2.0
// @description This is a sample server for Swagger.
// @host localhost:8070
// @BasePath /api/v1
func GinRun(config util.Config) {

	server, err := server.NewServer(config)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}
	server.Start(config.ServerAddress)
}

func GrpcRun(config util.Config) {
	server, err := gapi.NewServer(config)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterSocialMediaFoodServer(grpcServer, server)

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GrpcAddress)
	if err != nil {
		log.Fatal("cannot make grpc server:", err)
	}
	log.Printf("Start grpc server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot make grpc server:", err)
	}
}
