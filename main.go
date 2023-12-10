package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thomaszub/go-todos-templ-htmx/controller"
)

//go:embed assets
var assets embed.FS

func main() {
	c := controller.NewTodos()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", c.Get)
	r.Route("/assets", func(r chi.Router) {
		r.Get("/*", http.FileServer(http.FS(assets)).ServeHTTP)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
