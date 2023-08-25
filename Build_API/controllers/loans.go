package controllers

import (
	"Build_API/models"
	"Build_API/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func create customer record  --- CREATE
func CreateLoanRecord(c *gin.Context) {
	var temp models.Loans
	if err := c.ShouldBindJSON(&temp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cusID := temp.Customer_Id
	check := services.Loan_Check_Id(cusID, Collection)

	if !check {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Create Customer Record to Create Loans"})
	} else {
		services.CreateLoanRecord(temp, Collection2)
		c.JSON(http.StatusOK, gin.H{"message": "Loan record created"})
	}
}

// GET ALL RECORD ---- READ
func GetAllLoanRecord(c *gin.Context) {
	allRecord := services.GetAllLoanRecord(Collection2)
	c.JSON(http.StatusOK, allRecord)
}

// Update Loan Amount --- UPDATE
func UpdateLoanRecord(c *gin.Context) {
	id := c.Param("id")
	services.UpdateLoanAmount(id, Collection2)
	c.JSON(http.StatusOK, gin.H{"message": "Loan Amount updated"})
}

// // DeleteCustomerRecord --- DELETE
func DeleteLoanRecord(c *gin.Context) {
	id := c.Param("id")
	services.DeleteLoanRecord(id, Collection2)
	c.JSON(http.StatusOK, gin.H{"message": "Loan record deleted"})
}
