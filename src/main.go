package main

import (
	database "gin-ecom-backend/config"
	"gin-ecom-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main(){
	database.Connect()
	database.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
		&models.ShippingAddress{},
	)
	r := gin.Default()

	r.GET("/health", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
}