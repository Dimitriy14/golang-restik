package drinks

import (
	"net/http"

	"github.com/Dimitriy14/golang-restik/src/models"
	"github.com/google/uuid"

	"github.com/Dimitriy14/golang-restik/src/dal"
	"github.com/Dimitriy14/golang-restik/src/logger"
	"github.com/Dimitriy14/golang-restik/src/services/common"
)

type Service interface {
	GetDrinks(w http.ResponseWriter, r *http.Request)
	SaveDrink(w http.ResponseWriter, r *http.Request)
}

type drinksImpl struct {
	repo dal.Repository
	log  logger.Logger
}

func NewService(repo dal.Repository, log logger.Logger) Service {
	return &drinksImpl{
		repo: repo,
		log:  log,
	}
}

func (d *drinksImpl) GetDrinks(w http.ResponseWriter, r *http.Request) {
	drinks, err := d.repo.GetDrinks(r.Context())
	if err != nil {
		d.log.Errorf("", "cannot get drinks: %s", err)
		common.SendInternalServerError(w, "cannot find any drinks", err)
	}

	common.RenderJSON(w, drinks)
}

func (d *drinksImpl) SaveDrink(w http.ResponseWriter, r *http.Request) {
	drink := models.Drink{
		ID:     uuid.New(),
		Name:   "MyDrink",
		Volume: 0.5,
		Price:  115,
	}
	err := d.repo.AddNewDrink(r.Context(), drink)

	if err != nil {
		d.log.Errorf("", "cannot get drinks: %s", err)
		common.SendInternalServerError(w, "cannot find any drinks", err)
	}

	common.RenderJSONCreated(w, drink)
}
