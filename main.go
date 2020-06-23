package main

import (
	"companyservice/messaging"
	handlers2 "companyservice/messaging/handlers"
	"companyservice/routers"
	"companyservice/utils"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var (
	g errgroup.Group
)

func main() {
	router := routers.InitRoute()
	port := utils.EnvVar("SERVER_PORT")

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	handlers := map[string]messaging.MessageHandler{}
	handlers["UserChanged"] = handlers2.UserChangedHandler{}

	err := messaging.AddMessageConsumer(utils.EnvVar("RABBITMQ_CONNECTION_STRING"), "companyservice", handlers)
	log.Print("[AMQP] Started AMQP server")
	if err != nil {
		log.Fatal(err)
	}

	g.Go(func() error {
		log.Print("[HTTP] Started HTTP server")
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
