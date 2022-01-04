package route_handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (app *App) Movie(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		str := []byte(`{"Error" : "id not a integer"}`)
		w.Write(str)
		return
	}

	movie, err := app.DBhandler.ReadWithId(id)

	if err != nil {
		str := []byte(`{"Error" : "404 not Found"}`)
		w.Write(str)
		return
	}

	str, _ := json.Marshal(movie)
	w.Write(str)
}

func (app *App) AddMovie(w http.ResponseWriter, r *http.Request) {
	// Handle Write Movie
}
