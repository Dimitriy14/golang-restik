package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Food struct {
	ID           uuid.UUID     `json:"id"          gorm:"column:id"`
	Name         string        `json:"name"        gorm:"column:name"`
	Type         string        `json:"type"        gorm:"column:type"`
	Description  string        `json:"description" gorm:"column:description"`
	Image        string        `json:"image"       gorm:"column:image"`
	Price        float64       `json:"price"       gorm:"column:price"`
	Reservations []Reservation `json:"-" gorm:"AssociationForeignKey:ID; ForeignKey:ID"`
}

func (d *Food) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", uuid.New())
}
