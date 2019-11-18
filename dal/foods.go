package dal

import (
	"context"

	"github.com/Dimitriy14/golang-restik/models"
)

func (d *repoImpl) GetFoods(ctx context.Context) ([]models.Food, error) {
	var foods []models.Food

	if err := d.db.Session.Find(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}

func (d *repoImpl) SaveFood(ctx context.Context, food models.Food) (models.Food, error) {
	if err := d.db.Session.Save(&food).Error; err != nil {
		return models.Food{}, err
	}
	return food, nil
}
