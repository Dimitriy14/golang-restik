package services

import (
	"net/http"
	"strings"

	"github.com/Dimitriy14/golang-restik/services/foods"

	"github.com/Dimitriy14/golang-restik/services/tables"

	"github.com/Dimitriy14/golang-restik/config"
	"github.com/Dimitriy14/golang-restik/dal"
	"github.com/Dimitriy14/golang-restik/logger"
	"github.com/Dimitriy14/golang-restik/services/reservations"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func NewRouter() http.Handler {
	repo := dal.NewRepository()
	rs := reservations.NewService(repo, logger.Log)
	ts := tables.NewService(repo, logger.Log)
	foodServ := foods.NewService(repo, logger.Log)

	router := mux.NewRouter().StrictSlash(true).PathPrefix(config.Conf.BasePath).Subrouter()

	router.HandleFunc("/tables", ts.GetTables).Methods(http.MethodGet)
	router.HandleFunc("/tables", ts.AddTable).Methods(http.MethodPost)

	router.HandleFunc("/reservations", rs.GetReservations).Methods(http.MethodGet)
	router.HandleFunc("/reservations", rs.SaveReservation).Methods(http.MethodPost)

	router.HandleFunc("/foods", foodServ.GetFoods).Methods(http.MethodGet)
	router.HandleFunc("/foods", foodServ.AddNewFood).Methods(http.MethodPost)

	corsRouter := mux.NewRouter()
	{
		corsRouter.PathPrefix(config.Conf.BasePath).Handler(negroni.New(
			cors.New(cors.Options{
				AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
			}),
			negroni.Wrap(router),
		))
	}

	return removeTrailingSlash(corsRouter)
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}
