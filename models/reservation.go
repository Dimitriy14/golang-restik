package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Reservation struct {
	ID               uuid.UUID         `json:"id"      gorm:"column:id"`
	Phone            string            `json:"phone"   gorm:"column:phone"`
	Name             string            `json:"name"    gorm:"column:name"`
	StartAt          time.Time         `json:"startAt" gorm:"column:start_at"`
	EndAt            time.Time         `json:"endAt"   gorm:"column:end_at"`
	Table            uuid.UUID         `json:"tableID" gorm:"column:table_id"`
	ReservationFoods []ReservationFood `json:"foods"   gorm:"AssociationForeignKey:ID; ForeignKey:ReservationID"`
}

type ReservationFood struct {
	ReservationID uuid.UUID `json:"reservationID" gorm:"column:reservation_id"`
	FoodID        uuid.UUID `json:"foodID" gorm:"column:food_id"`
	Amount        int       `json:"amount" gorm:"column:amount"`
}

func (r ReservationFood) TableName() string {
	return "reserved_food"
}

func (d *Reservation) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.New())
}

type ReservedTime struct {
	ID      uuid.UUID `json:"id"      gorm:"column:id"`
	StartAt time.Time `json:"startAt" gorm:"column:start_at"`
	EndAt   time.Time `json:"endAt"   gorm:"column:end_at"`
	Table   uuid.UUID `json:"tableID" gorm:"column:table_id"`
}

func (ReservedTime) TableName() string {
	return "reservations"
}
