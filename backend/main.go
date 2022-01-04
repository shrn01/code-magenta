package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/shrn01/code-scarlet/db_handlers"
	"github.com/shrn01/code-scarlet/route_handlers"
)

func main() {
	r := mux.NewRouter()
	app := &route_handlers.App{
		DBhandler: db_handlers.GetHandler(),
	}

	api := r.PathPrefix("/api").Subrouter()

	fmt.Println("Started")

	api.HandleFunc("/movies", app.Movies).Methods("GET")
	api.HandleFunc("/movie/{id}", app.Movie).Methods("GET")
	api.HandleFunc("/movie", app.AddMovie).Methods("POST")

	api.HandleFunc("/like/{id}", app.HandleLikes).Methods("POST")
	api.HandleFunc("/genres", app.Genres)

	origins := handlers.AllowedOrigins([]string{"*"})

	http.ListenAndServe(":8000", handlers.CORS(origins)(r))
}
