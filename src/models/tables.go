package models

import (
	"github.com/google/uuid"
)

type Table struct {
	ID           uuid.UUID `json:"id"`
	Orders       []Order   `json:"orders"`
	NowAvailable bool      `json:"nowAvailable"`
}
