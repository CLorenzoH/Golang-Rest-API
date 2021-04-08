package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

func all_articles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{
			Title:   "title article",
			Desc:    "description of the article",
			Content: "content of the article",
		},
	}

	fmt.Println("All articles endpoint hit")
	json.NewEncoder(w).Encode(articles)
}

func post_articles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test post article endpoint hit")
}

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test homepage endpoint hit")
}

func handle_requests() {
	my_router := mux.NewRouter().StrictSlash(true)

	my_router.HandleFunc("/", home_page)

	my_router.HandleFunc("/articles", all_articles).Methods("GET")
	my_router.HandleFunc("/articles", post_articles).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", my_router))
}

func main() {
	handle_requests()
}
