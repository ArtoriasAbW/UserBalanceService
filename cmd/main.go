package main

import (
	balance "UserBalanceService"
	"UserBalanceService/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(balance.Server)
	err := srv.Run("8000", handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}