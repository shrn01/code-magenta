package route_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shrn01/code-scarlet/models"
)

func (app *App) Movie(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	movie := models.Movie{}

	result := app.DBhandler.Db.First(&movie, id)

	if result.Error != nil {
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
