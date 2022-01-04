package db_handlers

import (
	"github.com/shrn01/code-scarlet/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func GetHandler() Handler {
	handler := Handler{}
	if MODE == "dev" {
		handler.db, _ = gorm.Open(sqlite.Open(URL), &gorm.Config{})
	} else {
		handler.db, _ = gorm.Open(postgres.Open(URL), &gorm.Config{})
	}

	return handler
}

func (h *Handler) Create(movie models.Movie) interface{} {
	result := h.db.Create(movie)
	return result
}

func (h *Handler) ReadWithId(id int) (models.Movie, error) {
	var movie models.Movie
	result := h.db.First(&movie, id)
	return movie, result.Error
}

func (h *Handler) Readall() ([]map[string]interface{}, error) {
	var movies []map[string]interface{}
	result := h.db.Model(&models.Movie{}).Select("id, movie").Find(&movies)
	return movies, result.Error
}

// Db.Model(&models.Movie{}).Select("id, movie").Find(&movies)

func (h *Handler) Update(movie models.Movie) interface{} {
	result := h.db.Save(movie)
	return result
}

func (h *Handler) Delete(movie models.Movie) bool {
	result := h.db.Delete(movie)
	return result.Error == nil
}
