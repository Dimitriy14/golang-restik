package models

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/google/uuid"
)

type Reservations struct {
	ID      uuid.UUID `json:"id"`
	Visitor string    `json:"visitor"`
	StartAt time.Time `json:"startAt"`
	EndAt   time.Time `json:"endAt"`
	Drinks  []Drink   `json:"drinks" gorm:"many2many:reserved_drinks;"`
}

func (d *Reservations) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.New())
}
