package db_handlers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Handler struct {
	DbUrl string
	Db    *gorm.DB
}

func (h *Handler) Create() {
	h.Db, _ = gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
}
