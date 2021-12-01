package main

import (
	balance "UserBalanceService"
	"UserBalanceService/pkg/handler"
	"UserBalanceService/pkg/repository"
	"UserBalanceService/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(balance.Server)
	err := srv.Run("8000", handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}