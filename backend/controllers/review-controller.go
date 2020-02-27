package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ReviewsHandler(r *mux.Router) {
	r.HandleFunc("", fetchAllReviews)
	r.HandleFunc("/{reviewId}", fetchReviewById)
}

func fetchAllReviews(w http.ResponseWriter, r *http.Request) {
	fmt.Println("from handler")
	w.Write([]byte("reviews"))
}

func fetchReviewById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reviewId := vars["reviewId"]
	w.Write([]byte(reviewId))
}
