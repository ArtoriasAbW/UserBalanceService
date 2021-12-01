package handler

import (
	balance "UserBalanceService"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


func (h *Handler) refill(c *gin.Context){
	var input balance.Operation
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := h.services.Balance.IncreaseBalance(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("balance of user %d changed", userId),
	})
}


func (h *Handler) debit(c *gin.Context){

}


func (h *Handler) transfer(c *gin.Context){

}

// в логике всех хенделеров выше мы должны получать получить пользователей из бд
// если их нет выдать ошибку
// если есть посчитать новое значение баланса
// или выдать сообщение о невозможности изменить

func (h *Handler) balance(c *gin.Context) {
	var input balance.User
	if err := c.BindUri(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.services.Balance.GetUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": user.Id,
		"balance": user.Balance,
	})
}

