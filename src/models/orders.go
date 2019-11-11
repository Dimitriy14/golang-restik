package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID           uuid.UUID
	Visitor      string
	StartBooking time.Time
	EndBooking   time.Time
	Drinks       []Drink
}
