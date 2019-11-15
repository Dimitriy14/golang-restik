package dal

import (
	"context"
	"fmt"

	"github.com/Dimitriy14/golang-restik/models"
	"github.com/google/uuid"
)

func (r *repoImpl) GetDrinks(ctx context.Context) ([]models.Drink, error) {
	var drinks []models.Drink

	err := r.db.Session.Find(&drinks).Error
	if err != nil {
		return nil, err
	}

	return drinks, nil
}

func (r *repoImpl) GetDrinkByID(ctx context.Context, drinkID uuid.UUID) (models.Drink, error) {
	var drink = models.Drink{ID: drinkID}

	err := r.db.Session.Find(&drink).Error
	if err != nil {
		return models.Drink{}, err
	}

	return drink, nil
}

func (r *repoImpl) AddNewDrink(ctx context.Context, drink models.Drink) (models.Drink, error) {
	errs := r.db.Session.Create(&drink).GetErrors()
	if len(errs) > 1 {
		return models.Drink{}, fmt.Errorf("cannot save drink: %v", errs)
	}

	return drink, nil
}

func (r *repoImpl) UpdateDrink(ctx context.Context, drink models.Drink) (models.Drink, error) {
	if err := r.db.Session.Model(&models.Drink{}).Updates(drink).Error; err != nil {
		return models.Drink{}, err
	}
	return drink, nil
}

func (r *repoImpl) DeleteDrink(ctx context.Context, drinkID uuid.UUID) error {
	errs := r.db.Session.Delete(&models.Drink{ID: drinkID}).GetErrors()
	if len(errs) > 1 {
		return fmt.Errorf("cannot save drink: %v", errs)
	}

	return nil
}
