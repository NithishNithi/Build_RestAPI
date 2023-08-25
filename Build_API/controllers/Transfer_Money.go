package controllers

import (
	"Build_API/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeTransaction(c *gin.Context) {
	var trans_money struct {
		From_Id int `json:"from_id"`
		To_Id   int `json:"to_Id"`
		Amount  int `json:"amount"`
	}
	if err := c.ShouldBindJSON(&trans_money); err != nil {
		fmt.Println("Bind error")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cus1_balance, cus2_balance := services.TransferMoney(trans_money.From_Id, trans_money.To_Id, trans_money.Amount, Collection, Collection3)
	c.JSON(http.StatusOK, cus1_balance)
	c.JSON(http.StatusOK, cus2_balance)

}
