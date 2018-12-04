package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/render"
)

func main() {
	r := chi.NewRouter()
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		res := map[string]string{
			"message": "pong",
		}
		render.JSON(w, r, res)
	})
	http.ListenAndServe(":3000", r)
}
