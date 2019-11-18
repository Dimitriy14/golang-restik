package reservations

import (
	"encoding/json"
	"net/http"

	"github.com/Dimitriy14/golang-restik/models"

	"github.com/Dimitriy14/golang-restik/services/common"

	"github.com/Dimitriy14/golang-restik/dal"
	"github.com/Dimitriy14/golang-restik/logger"
	txID "github.com/Dimitriy14/golang-restik/logger/transaction-id"
)

type Service interface {
	GetReservations(w http.ResponseWriter, r *http.Request)
	SaveReservation(w http.ResponseWriter, r *http.Request)
}

type reservationServiceImpl struct {
	repo dal.Repository
	log  logger.Logger
}

func NewService(repo dal.Repository, log logger.Logger) Service {
	return &reservationServiceImpl{
		repo: repo,
		log:  log,
	}
}

func (d *reservationServiceImpl) GetReservations(w http.ResponseWriter, r *http.Request) {
	txid := txID.FromRequest(r)

	d.log.Debug(txid, "Started retrieving reservations")

	drinks, err := d.repo.GetReservations(r.Context())
	if err != nil {
		d.log.Errorf(txid, "cannot get reservations: %s", err)
		common.SendInternalServerError(w, "cannot find any reservations", err)
		return
	}

	d.log.Debug(txid, "Successfully retrieved reservations")

	common.RenderJSON(w, drinks)
}

func (d *reservationServiceImpl) SaveReservation(w http.ResponseWriter, r *http.Request) {
	txid := txID.FromRequest(r)

	d.log.Debug(txid, "Started saving reservations")

	var reservation models.Reservation

	err := json.NewDecoder(r.Body).Decode(&reservation)
	if err != nil {
		d.log.Errorf(txid, "cannot decode reservation: %s", err)
		common.SendInternalServerError(w, "cannot decode reservation", err)
		return
	}
	defer common.CloseReqBody(r)

	reservation, err = d.repo.SaveReservation(r.Context(), reservation)
	if err != nil {
		d.log.Errorf(txid, "cannot save reservation: %s", err)
		common.SendInternalServerError(w, "cannot save reservation", err)
		return
	}

	d.log.Debug(txid, "Successfully saved reservations")

	common.RenderJSONCreated(w, reservation)
}
