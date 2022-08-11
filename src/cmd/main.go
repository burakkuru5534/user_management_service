package main

import (
	"example.com/m/v2/src/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	_ "github.com/lib/pq"
)

func main() {

	//router
	r := mux.NewRouter()
	//api endpoints
	r.Handle("/users", api.UserCreate())
	r.Handle("/users/{id}", api.UserUpdate())
	r.Handle("/users/{id}", api.UserGet())
	r.Handle("/users", api.UserList())
	r.Handle("/users/{id}", api.UserDelete())

	//define options
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	//start server
	log.Fatal(http.ListenAndServe(":8080", corsWrapper.Handler(r)))
}
