package main

import (
	"fmt"
	"log"
	"main/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	controller := controllers.NewController()
	controller.RegisterArticleController(router)

	port := ":8080"
	fmt.Printf("Connected to %s\n", port)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
