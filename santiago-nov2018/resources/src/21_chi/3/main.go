package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/render"
	"github.com/ifreddyrondon/go-workshop/santiago-nov2018/resources/src/21_chi/3/todo"
)

func routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,                             // Log API request calls
		middleware.Recoverer,                          // Recover from panics without crashing server
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/todo", todo.Routes())
	})

	return router
}

func main() {
	//db := gorm.Open("postgres", "postgres://localhost/go_workshop?sslmode=disable")

	log.Fatal(http.ListenAndServe(":3000", routes())) // Note, the port is usually gotten from the environment.
}
