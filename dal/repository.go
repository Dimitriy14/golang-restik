package dal

import (
	"context"

	"github.com/Dimitriy14/golang-restik/clients/postgres"
	"github.com/Dimitriy14/golang-restik/models"
)

type Repository interface {
	GetTables(ctx context.Context) ([]models.Table, error)
	AddTable(ctx context.Context, table models.Table) (models.Table, error)

	GetReservations(ctx context.Context) ([]models.Reservation, error)
	SaveReservation(ctx context.Context, reservation models.Reservation) (models.Reservation, error)

	GetFoods(ctx context.Context) ([]models.Food, error)
	SaveFood(ctx context.Context, food models.Food) (models.Food, error)
}

type repoImpl struct {
	db postgres.PGClient
}

func NewRepository() Repository {
	return &repoImpl{db: postgres.Client}
}
