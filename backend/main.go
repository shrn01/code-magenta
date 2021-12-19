package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	routehandlers "github.com/shrn01/code-scarlet/route_handlers"
)

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	fmt.Println("Started")

	api.HandleFunc("/", routehandlers.Home)
	api.HandleFunc("/movie", routehandlers.Movie)
	api.HandleFunc("/like", routehandlers.HandleLikes)

	http.ListenAndServe(":8000", r)
}
