package main

import (
	"Build_API/routers"
	"fmt"
	"log"
	"net/http"
)

func main(){
	fmt.Println("Build-API MangoDB ")
	fmt.Println("server gets Started")
	r:=routers.Routers()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening and serving on Port:4000")
	
}