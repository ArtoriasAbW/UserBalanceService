package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {

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
