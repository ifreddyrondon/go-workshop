package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func handler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	w.Write([]byte(fmt.Sprintf("hi %v", userID)))
}

func main() {
	r := chi.NewRouter()
	r.Get("/user/:userID", handler)
	http.ListenAndServe(":3000", r)
}
