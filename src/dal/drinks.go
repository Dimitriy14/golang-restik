package dal

import (
	"context"
	"fmt"

	"github.com/Dimitriy14/golang-restik/src/models"
)

func (r *repoImpl) GetDrinks(ctx context.Context) ([]models.Drink, error) {
	var drinks []models.Drink

	rows, err := r.db.Session.Table("drinks").Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var drink models.Drink

		if err = rows.Scan(&drink.ID, &drink.Name, &drink.Price, &drink.Volume); err != nil {
			return nil, err
		}

		drinks = append(drinks, drink)
	}
	return drinks, nil
}

func (r *repoImpl) AddNewDrink(ctx context.Context, drink models.Drink) error {
	errs := r.db.Session.Table("drinks").Save(&drink).GetErrors()
	if len(errs) > 1 {
		return fmt.Errorf("cannot save drink: %v", errs)
	}

	return nil
}
