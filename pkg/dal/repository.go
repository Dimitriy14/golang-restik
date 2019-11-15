package dal

import (
	"context"

	"github.com/Dimitriy14/golang-restik/pkg/clients/postgres"
	"github.com/Dimitriy14/golang-restik/pkg/models"
	"github.com/google/uuid"
)

type Repository interface {
	GetDrinks(ctx context.Context) ([]models.Drink, error)
	GetDrinkByID(ctx context.Context, drinkID uuid.UUID) (models.Drink, error)
	AddNewDrink(ctx context.Context, drink models.Drink) (models.Drink, error)
	UpdateDrink(ctx context.Context, drink models.Drink) (models.Drink, error)
	DeleteDrink(ctx context.Context, drinkID uuid.UUID) error
}

type repoImpl struct {
	db postgres.PGClient
}

func NewRepository() Repository {
	return &repoImpl{db: postgres.Client}
}
