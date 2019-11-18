package dal

import (
	"context"
	"fmt"

	"github.com/Dimitriy14/golang-restik/models"
)

func (r *repoImpl) GetTables(ctx context.Context) ([]models.Table, error) {
	var tables []models.Table

	err := r.db.Session.Preload("Reservations").Find(&tables).Error
	if err != nil {
		return nil, err
	}

	return tables, nil
}

func (r *repoImpl) GetReservations(ctx context.Context) ([]models.Reservation, error) {
	var reservations []models.Reservation

	err := r.db.Session.Find(&reservations).Error
	if err != nil {
		return nil, err
	}

	return reservations, nil
}

func (r *repoImpl) SaveReservation(ctx context.Context, reservation models.Reservation) (models.Reservation, error) {
	if !r.checkIfFoodExist(reservation.ReservationFoods...) {
		return models.Reservation{}, fmt.Errorf("some food doesn't exist %v", reservation.ReservationFoods)
	}

	err := r.db.Session.Save(&reservation).Error
	if err != nil {
		return models.Reservation{}, err
	}
	return reservation, nil
}

func (r *repoImpl) AddTable(ctx context.Context, table models.Table) (models.Table, error) {
	err := r.db.Session.Save(&table).Error
	if err != nil {
		return table, err
	}

	return table, err
}

func (r *repoImpl) checkIfFoodExist(foods ...models.ReservationFood) bool {
	for _, food := range foods {
		if r.db.Session.Find(&models.Food{ID: food.FoodID}).RecordNotFound() {
			return false
		}
	}
	return true
}
