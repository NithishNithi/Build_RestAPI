package main

import (
	"Build_API/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Build-API MongoDB")
	fmt.Println("Server gets started")
	r := routers.Routers()
	fmt.Println("Listening and serving on Port:4000")
	err := http.ListenAndServe(":4000", r)
	if err != nil {
		log.Println("Error:", err)
	}
}
