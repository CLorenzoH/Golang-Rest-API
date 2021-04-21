package controllers

import (
	"encoding/json"
	"main/models"
	"net/http"
)

func (ControllerHandler) GetArticles(w http.ResponseWriter, r *http.Request) {
	articles := []models.Article{
		{
			Title:       "title article",
			Description: "description of the article",
			Content:     "content of the article",
		},
	}

	w.Header().Add("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(articles)
}

func (ControllerHandler) PostArticles(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Endpoint post article hit"))
}
