package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Drink struct {
	ID     uuid.UUID `json:"id" sql:"id"`
	Name   string    `json:"name" sql:"name"`
	Volume float64   `json:"volume" sql:"volume"`
	Price  float64   `json:"price" sql:"price"`
}

func (d *Drink) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", uuid.New())
}
