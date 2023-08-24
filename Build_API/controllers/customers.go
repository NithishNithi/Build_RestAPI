package controllers

import (
	"Build_API/models"
	"Build_API/password"
	"Build_API/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// func create customer record  --- CREATE
func CreateCustomerRecord(c *gin.Context) {
	c.String(http.StatusOK, "Inserted in Postman")
	var temp models.Customers
	if err := c.ShouldBindJSON(&temp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := password.HashPassword(temp.Password) // hashing the password before inserting into database
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	temp.Password = hashedPassword

	services.CreateCustomerRecord(temp, Collection)
	c.JSON(http.StatusOK, temp)
}

// GET ALL RECORD ---- READ
func GetAllCustomerRecord(c *gin.Context) {
	allRecord := services.GetAllCustomerRecord(Collection)
	c.JSON(http.StatusOK, allRecord)
}

// UpdateCustomerPassword --- UPDATE
func UpdateCustomerPassword(c *gin.Context) {
	id := c.Param("id")
	services.UpdateCustomerPassword(id, Collection)
	c.JSON(http.StatusOK, gin.H{"message": "Password updated"})
}

// DeleteCustomerRecord --- DELETE
func DeleteCustomerRecord(c *gin.Context) {
	id := c.Param("id")
	services.DeleteCustomerRecord(id, Collection)
	c.JSON(http.StatusOK, gin.H{"message": "Customer record deleted"})
}
