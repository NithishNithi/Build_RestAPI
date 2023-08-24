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
	port:=":4002"
	fmt.Println("Listening and serving on Port",port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Println("Error:", err)
	}
}
