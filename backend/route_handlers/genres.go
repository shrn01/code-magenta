package route_handlers

import (
	"net/http"
)

func Genres(w http.ResponseWriter, r *http.Request) {
	bytes := []byte("not yet implemented")
	w.Write(bytes)
}
