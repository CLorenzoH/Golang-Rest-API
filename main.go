package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test homepage endpoint hit")
}

func handle_requests() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/articles", all_articles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handle_requests()
}
