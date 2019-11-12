package dal

import (
	"context"

	"github.com/Dimitriy14/golang-restik/src/clients/postgres"
	"github.com/Dimitriy14/golang-restik/src/models"
)

type Repository interface {
	GetDrinks(ctx context.Context) ([]models.Drink, error)
	AddNewDrink(ctx context.Context, drink models.Drink) error
}

type repoImpl struct {
	db postgres.PGClient
}

func NewRepository() Repository {
	return &repoImpl{db: postgres.Client}
}
