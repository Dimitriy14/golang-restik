package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Table struct {
	ID           uuid.UUID      `json:"id"`
	Orders       []Reservations `json:"reservations"`
	NowAvailable bool           `json:"nowAvailable"`
}

func (d *Table) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.New())
}
