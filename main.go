package main

import (
	"fmt"
	"log"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test homepage endpoint hit")
}

func handle_requests() {
	http.HandleFunc("/", home_page)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handle_requests()
}
