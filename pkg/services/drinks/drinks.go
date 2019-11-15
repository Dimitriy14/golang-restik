package drinks

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Dimitriy14/golang-restik/pkg/dal"
	"github.com/Dimitriy14/golang-restik/pkg/logger"
	txID "github.com/Dimitriy14/golang-restik/pkg/logger/transaction-id"
	"github.com/Dimitriy14/golang-restik/pkg/models"
	"github.com/Dimitriy14/golang-restik/pkg/services/common"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Service interface {
	GetDrinks(w http.ResponseWriter, r *http.Request)
	GetDrinkByID(w http.ResponseWriter, r *http.Request)
	SaveDrink(w http.ResponseWriter, r *http.Request)
	UpdateDrink(w http.ResponseWriter, r *http.Request)
	DeleteDrink(w http.ResponseWriter, r *http.Request)
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
	txid := txID.FromRequest(r)

	d.log.Debug(txid, "Started retrieving drink")

	drinks, err := d.repo.GetDrinks(r.Context())
	if err != nil {
		d.log.Errorf(txid, "cannot get drinks: %s", err)
		common.SendInternalServerError(w, "cannot find any drinks", err)
		return
	}

	d.log.Debug(txid, "Successfully retrieved drink")

	common.RenderJSON(w, drinks)
}

func (d *drinksImpl) GetDrinkByID(w http.ResponseWriter, r *http.Request) {
	var (
		txid = txID.FromRequest(r)
		id   = mux.Vars(r)["id"]
	)

	d.log.Debugf(txid, "Started retrieving drink with id: %s", id)

	uuid, err := uuid.Parse(id)
	if err != nil {
		d.log.Errorf(txid, "getting drink id: %s", err)
		common.SendError(w, http.StatusBadRequest, "invalid drink id", err)
		return
	}

	drink, err := d.repo.GetDrinkByID(r.Context(), uuid)
	if err != nil {
		d.log.Errorf(txid, "cannot get drinks: %s", err)
		common.SendInternalServerError(w, "cannot find any drinks", err)
		return
	}

	d.log.Debugf(txid, "Successfully retrieved a drink with id: %s", id)

	common.RenderJSON(w, drink)
}

func (d *drinksImpl) SaveDrink(w http.ResponseWriter, r *http.Request) {
	txid := txID.FromRequest(r)

	d.log.Debug(txid, "Started creating new drink")

	drink, err := getDrinkFromReq(r)
	if err != nil {
		common.SendError(w, http.StatusBadRequest, "Invalid body", err)
		return
	}

	drink, err = d.repo.AddNewDrink(r.Context(), drink)
	if err != nil {
		d.log.Errorf(txid, "got error while creating new drinks: %s", err)
		common.SendInternalServerError(w, "got error while creating new drinks", err)
		return
	}

	d.log.Debug(txid, "Successfully created new drink")

	common.RenderJSONCreated(w, drink)
}

func (d *drinksImpl) UpdateDrink(w http.ResponseWriter, r *http.Request) {
	var (
		txid = txID.FromRequest(r)
		id   = mux.Vars(r)["id"]
	)

	d.log.Debugf(txid, "Started updating drink with id: %s", id)

	uuid, err := uuid.Parse(id)
	if err != nil {
		d.log.Errorf(txid, "getting drink id: %s", err)
		common.SendError(w, http.StatusBadRequest, "invalid drink id", err)
		return
	}

	drink, err := getDrinkFromReq(r)
	if err != nil {
		d.log.Errorf(txid, "retrieving drink from body: %s", err)
		common.SendError(w, http.StatusBadRequest, "invalid drink", err)
		return
	}
	drink.ID = uuid

	drink, err = d.repo.UpdateDrink(r.Context(), drink)
	if err != nil {
		d.log.Errorf(txid, "updating drink error: %s", err)
		common.SendInternalServerError(w, "cannot update drink", err)
		return
	}

	d.log.Debugf(txid, "Successfully updated drink with id: %s", id)

	common.RenderJSON(w, drink)
}

func (d *drinksImpl) DeleteDrink(w http.ResponseWriter, r *http.Request) {
	var (
		txid = txID.FromRequest(r)
		id   = mux.Vars(r)["id"]
	)

	d.log.Debugf(txid, "Started deleting drink with id: %s", id)

	uuid, err := uuid.Parse(id)
	if err != nil {
		d.log.Errorf(txid, "getting drink id: %s", err)
		common.SendError(w, http.StatusBadRequest, "invalid drink id", err)
		return
	}

	err = d.repo.DeleteDrink(r.Context(), uuid)
	if err != nil {
		d.log.Errorf(txid, "cannot get drinks: %s", err)
		common.SendInternalServerError(w, "cannot find any drinks", err)
		return
	}

	d.log.Debugf(txid, "Successfully deleted drink with id: %s", id)

	common.RenderNoContentStatus(w)
}

func getDrinkFromReq(r *http.Request) (models.Drink, error) {
	var drink models.Drink

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return models.Drink{}, err
	}
	defer common.CloseReqBody(r)

	if err = json.Unmarshal(body, &drink); err != nil {
		return models.Drink{}, err
	}

	return drink, nil
}
