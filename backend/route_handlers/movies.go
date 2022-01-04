package route_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shrn01/code-scarlet/models"
)

var value map[string]interface{} = map[string]interface{}{
	"id": 1, "movie": "Titanic", "year": 1997, "imdb": 7.8,
	"genre": "Romance", "actors": "", "likes": 0, "dislikes": 0, "summary": "", "added_by": "Tasak", "movie_or_series": "movie", "trailer": "https://www.youtube.com/watch?v=cIJ8ma0kKtY",
}

func (app *App) Movies(w http.ResponseWriter, r *http.Request) {

	movies, err := app.DBhandler.Readall()

	moviez := models.Movies{MovieList: movies}

	if err != nil {
		str := []byte(`{"Error" : "404 not Found"}`)
		w.Write(str)
		return
	}

	str, _ := json.Marshal(moviez)
	w.Write(str)
}
