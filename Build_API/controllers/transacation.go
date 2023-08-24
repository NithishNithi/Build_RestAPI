package controllers

import (
	"Build_API/models"
	"Build_API/services"
	"fmt"

	// "fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateTransaction --- CREATE
func CreateTransaction(c *gin.Context) {
	var temp models.Transactions
	if err := c.ShouldBindJSON(&temp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cusID := temp.Customer_Id
	check := services.Transaction_Check_Id(cusID, Collection)

	temp.Transaction_Time = time.Now()
	if !check {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Create Customer Record to Create Transaction"})
	} else {
		services.CreateTransactionRecord(temp, Collection1)
		c.JSON(http.StatusOK, gin.H{"message": "Transaction created"})
	}
}

// UpdateTransactionAmount --- UPDATE
func UpdateTransactionAmount(c *gin.Context) {
	id := c.Param("id")
	services.UpdateTransactionAmount(id, Collection1)
	c.JSON(http.StatusOK, gin.H{"message": "Transaction amount updated"})
}

// GetAllTransaction -- READ

func GetAllTransaction(c *gin.Context) {
	allRecord := services.GetAllTransactionRecord(Collection1)
	c.JSON(http.StatusOK, allRecord)
}

//

func GetTransactionsByTimeRange(c *gin.Context) {

	var req struct {
		Fr string `json:"start_time" binding:"required"`
		To string `json:"end_time" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("Bind error")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(req.Fr, req.To)

	starttime, err := time.Parse("2006-01-02", req.Fr)
	fmt.Println(req.Fr)
	if err != nil {
		fmt.Println("err")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_time format"})
		return
	}
	endtime, err := time.Parse("2006-01-02", req.To)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_time format"})
		return
	}
	transactions, _ := services.GetTransactionsByTimeRange(starttime, endtime, Collection1)

	c.JSON(http.StatusOK,transactions,)
}

// GetTransactionsSumByTimeRange

func GetTransactionsSumByTimeRange(c *gin.Context){
	var reqsum struct {
		Fr string `json:"start_time" binding:"required"`
		To string `json:"end_time" binding:"required"`
	}
	if err := c.ShouldBindJSON(&reqsum); err != nil {
		fmt.Println("Bind error")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(reqsum.Fr, reqsum.To)

	starttime, err := time.Parse("2006-01-02", reqsum.Fr)
	fmt.Println(reqsum.Fr)
	if err != nil {
		fmt.Println("err")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_time format"})
		return
	}
	endtime, err := time.Parse("2006-01-02", reqsum.To)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_time format"})
		return
	}
	totalamount,_ := services.GetTransactionsSumByTimeRange(starttime, endtime, Collection1)
	fmt.Println("Total Amount in Given Time Range:",totalamount)
	c.JSON(http.StatusOK,totalamount)
}


