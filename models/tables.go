package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Table struct {
	ID           uuid.UUID      `json:"id"   gorm:"primary_key"`
	Sits         uint16         `json:"sits" gorm:"column:sits"`
	Reservations []ReservedTime `json:"reservations,omitempty" gorm:"ForeignKey:Table; AssociationForeignKey:id"`
}

func (d *Table) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.New())
}
