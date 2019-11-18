package foods

import (
	"encoding/json"
	"net/http"

	"github.com/Dimitriy14/golang-restik/dal"
	"github.com/Dimitriy14/golang-restik/logger"
	txID "github.com/Dimitriy14/golang-restik/logger/transaction-id"
	"github.com/Dimitriy14/golang-restik/models"
	"github.com/Dimitriy14/golang-restik/services/common"
)

type Service interface {
	GetFoods(w http.ResponseWriter, r *http.Request)
	AddNewFood(w http.ResponseWriter, r *http.Request)
}

type foodServiceImpl struct {
	repo dal.Repository
	log  logger.Logger
}

func NewService(repo dal.Repository, log logger.Logger) Service {
	return &foodServiceImpl{
		repo: repo,
		log:  log,
	}
}

func (d *foodServiceImpl) GetFoods(w http.ResponseWriter, r *http.Request) {
	txid := txID.FromRequest(r)

	d.log.Debug(txid, "Started retrieving foods")

	foods, err := d.repo.GetFoods(r.Context())
	if err != nil {
		d.log.Errorf(txid, "cannot get foods: %s", err)
		common.SendInternalServerError(w, "cannot find any foods", err)
		return
	}

	d.log.Debug(txid, "Successfully retrieved foods")

	common.RenderJSON(w, foods)
}

func (d *foodServiceImpl) AddNewFood(w http.ResponseWriter, r *http.Request) {
	txid := txID.FromRequest(r)

	d.log.Debug(txid, "Started saving new food")

	var food models.Food

	err := json.NewDecoder(r.Body).Decode(&food)
	if err != nil {
		d.log.Errorf(txid, "cannot decode food: %s", err)
		common.SendInternalServerError(w, "cannot decode food", err)
		return
	}
	defer common.CloseReqBody(r)

	food, err = d.repo.SaveFood(r.Context(), food)
	if err != nil {
		d.log.Errorf(txid, "cannot save food: %s", err)
		common.SendInternalServerError(w, "cannot save food", err)
		return
	}

	d.log.Debug(txid, "Successfully saved food")

	common.RenderJSONCreated(w, food)
}
