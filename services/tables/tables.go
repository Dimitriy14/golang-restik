package tables

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
	GetTables(w http.ResponseWriter, r *http.Request)
	AddTable(w http.ResponseWriter, r *http.Request)
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

func (d *drinksImpl) GetTables(w http.ResponseWriter, r *http.Request) {
	txid := txID.FromRequest(r)

	d.log.Debug(txid, "Started retrieving tables")

	drinks, err := d.repo.GetTables(r.Context())
	if err != nil {
		d.log.Errorf(txid, "cannot get drinks: %s", err)
		common.SendInternalServerError(w, "cannot find any drinks", err)
		return
	}

	d.log.Debug(txid, "Successfully retrieved tables")

	common.RenderJSON(w, drinks)
}

func (d *drinksImpl) AddTable(w http.ResponseWriter, r *http.Request) {
	txid := txID.FromRequest(r)

	d.log.Debug(txid, "Started saving reservations")

	var table models.Table

	err := json.NewDecoder(r.Body).Decode(&table)
	if err != nil {
		d.log.Errorf(txid, "cannot decode reservation: %s", err)
		common.SendInternalServerError(w, "cannot decode reservation", err)
		return
	}
	defer common.CloseReqBody(r)

	table, err = d.repo.AddTable(r.Context(), table)
	if err != nil {
		d.log.Errorf(txid, "cannot save reservation: %s", err)
		common.SendInternalServerError(w, "cannot save reservation", err)
		return
	}

	d.log.Debug(txid, "Successfully saved reservations")

	common.RenderJSONCreated(w, table)
}
