package handler

import (
	"net/http"

	"github.com/codeedu/go-hexagonal/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeProductHandlers(router *mux.Router, negroni *negroni.Negroni, service application.ProductServiceInterface) {
	router.HandleFunc("/product/{id}", negroni.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
}