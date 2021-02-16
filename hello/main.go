package main

import (
	"hello/handler"
	pb "hello/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("hello"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterHelloHandler(srv.Server(), new(handler.Hello))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
