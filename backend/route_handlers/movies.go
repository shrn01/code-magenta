package route_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shrn01/code-scarlet/models"
)

func (app *App) Movies(w http.ResponseWriter, r *http.Request) {
	var movies []map[string]interface{}

	result := app.DBhandler.Db.Model(&models.Movie{}).Select("id, movie").Find(&movies)

	moviez := models.Movies{MovieList: movies}

	if result.Error != nil {
		str := []byte(`{"Error" : "404 not Found"}`)
		w.Write(str)
		return
	}

	str, _ := json.Marshal(moviez)
	w.Write(str)
}
