package handler

import (
	"UserBalanceService/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine  {
	router := gin.New()
	balanceService := router.Group("/balance-service")
	{
		balanceService.POST("/refill", h.refill)
		balanceService.POST("/debit", h.debit)
		balanceService.POST("/transfer", h.transfer)
		balanceService.GET("/balance/:id", h.balance)
	}
	return router
}
