package handler

import (
	balance "UserBalanceService"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) refill(c *gin.Context){

}


func (h *Handler) debit(c *gin.Context){

}


func (h *Handler) transfer(c *gin.Context){

}


func (h *Handler) balance(c *gin.Context){
	var input balance.User
	if err := c.BindUri(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.services.Balance.GetUser(input);
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": user.Id,
		"balance": user.Balance,
	})
}

