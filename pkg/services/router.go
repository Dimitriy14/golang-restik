package services

import (
	"net/http"
	"strings"

	"github.com/Dimitriy14/golang-restik/pkg/config"
	"github.com/Dimitriy14/golang-restik/pkg/dal"
	"github.com/Dimitriy14/golang-restik/pkg/logger"
	"github.com/Dimitriy14/golang-restik/pkg/services/drinks"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func NewRouter() http.Handler {
	repo := dal.NewRepository()
	ds := drinks.NewService(repo, logger.Log)

	router := mux.NewRouter().StrictSlash(true).PathPrefix(config.Conf.BasePath).Subrouter()

	router.HandleFunc("/drinks", ds.GetDrinks).Methods(http.MethodGet)
	router.HandleFunc("/drinks", ds.SaveDrink).Methods(http.MethodPost)
	router.HandleFunc("/drinks/{id}", ds.GetDrinkByID).Methods(http.MethodGet)
	router.HandleFunc("/drinks/{id}", ds.UpdateDrink).Methods(http.MethodPut)
	router.HandleFunc("/drinks/{id}", ds.DeleteDrink).Methods(http.MethodDelete)

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
