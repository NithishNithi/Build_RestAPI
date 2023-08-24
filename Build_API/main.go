package main

import (
	"Build_API/routers"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Build-API MongoDB")
	fmt.Println("Server gets started")
	r := gin.Default() //Create Gin router Instance
	routers.SetUpRouters(r)
	fmt.Println("Listening and serving on Port:4000")
	err := http.ListenAndServe(":4002", r)
	if err != nil {
		log.Println("Error:", err)
	}
}
