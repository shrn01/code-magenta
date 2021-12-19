package db_handlers

import (
	"github.com/shrn01/code-scarlet/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Handler struct {
	Db *gorm.DB
}

func GetHandler() Handler {
	handler := Handler{}
	if MODE == "dev" {
		handler.Db, _ = gorm.Open(sqlite.Open(URL), &gorm.Config{})
	} else {
		handler.Db, _ = gorm.Open(postgres.Open(URL), &gorm.Config{})
	}

	return handler
}

func (h *Handler) Create(movie models.Movie) {
	h.Db.Create(movie)
}

func (h *Handler) Read() {
	/*
		Implement this method to get movie object based on ID
	*/
}

func (h *Handler) Update(movie models.Movie) {
	h.Db.Save(movie)
}

func (h *Handler) Delete(movie models.Movie) {
	h.Db.Delete(movie)
}
