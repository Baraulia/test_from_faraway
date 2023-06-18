package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	client_handlers "test_faraway/quoteClient/handlers"
	client_services "test_faraway/quoteClient/services"
	"test_faraway/quoteServer/server"
)

func main() {
	service := client_services.Service{}
	handler := client_handlers.NewHandler(&service)

	apiServer := new(server.Server)

	if err := apiServer.Run("0.0.0.0", "8082", handler.InitRoutes()); err != nil {
		log.Panicf("Error occurred while running http server: %s", err.Error())
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("Server down")

	if err := apiServer.Shutdown(context.Background()); err != nil {
		log.Printf("error occurred on server shutting down: %s", err.Error())
	}

}
