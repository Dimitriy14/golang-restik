package models

import (
	"github.com/twinj/uuid"
	"time"
)

type Order struct {
	ID uuid.UUID
	Visitor string
	StartBooking time.Time
	EndBooking time.Time
	Drinks []Drink
}
