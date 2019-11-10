package models

import "github.com/twinj/uuid"

type Drink struct {
	ID uuid.UUID
	Name string
	Volume float64
	Price float64
}
