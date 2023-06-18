package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"test_faraway/quoteServer/handlers"
	"test_faraway/quoteServer/server"
	"test_faraway/quoteServer/services"
)

const (
	link   = "http://api.forismatic.com/api/1.0/"
	format = "json"
	method = "getQuote"
	lang   = "ru"
)

func main() {
	quoteService := services.NewQuoteService(link, method, format, lang)
	handler := handlers.NewHandler(quoteService)

	apiServer := new(server.Server)

	if err := apiServer.Run("0.0.0.0", "8081", handler.InitRoutes()); err != nil {
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
