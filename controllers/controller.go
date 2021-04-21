package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type ControllerHandler struct{}

func NewController() ControllerHandler {
	return ControllerHandler{}
}

func (c ControllerHandler) RegisterArticleController(r *mux.Router) {
	r.HandleFunc("/articles", c.GetArticles).Methods(http.MethodGet)
	r.HandleFunc("/articles", c.PostArticles).Methods(http.MethodPost)
}
