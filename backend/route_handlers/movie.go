package route_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/shrn01/code-scarlet/db_handlers"
	"github.com/shrn01/code-scarlet/models"
)

func Movie(w http.ResponseWriter, r *http.Request) {
	movie := models.Movie{}
	handler := db_handlers.GetHandler()

	handler.Db.First(&movie)

	spew.Dump(movie)

	str, _ := json.Marshal(movie)
	w.Write(str)
}
