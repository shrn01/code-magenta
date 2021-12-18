package route_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shrn01/code-scarlet/models"
)

func Movie(w http.ResponseWriter, r *http.Request) {
	movie := models.Movie{}

	str, _ := json.Marshal(movie)
	w.Write(str)
}
