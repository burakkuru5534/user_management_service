package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"net/http"
)

func HttpService() http.Handler {
	mux := chi.NewRouter()

	acors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	mux.Use(acors.Handler)

	mux.Route("/api", func(mr chi.Router) {
		mr.Group(func(r chi.Router) {
			//r.Get("/users/{id}", UserGet)
			r.Post("/users", UserCreate)
			r.Patch("/users", UserUpdate)
			r.Get("/users", UserList)
			r.Get("/user", UserGet)
			r.Delete("/users", UserDelete)
		})
	})

	return mux
}
