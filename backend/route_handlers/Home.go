package route_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shrn01/code-scarlet/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	movies := models.Movies{
		MovieIdList: []int{
			10,
			20,
			30,
		},
	}

	str, _ := json.Marshal(movies)
	w.Write(str)
}
