package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rafaelsanzio/go-consuming-soap/pkg/api/handlers"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()

	for _, r := range routes {
		router.Methods(r.Methods...).Path(r.Path).Name(r.Name).HandlerFunc(r.Handler)
	}

	return router
}

type Route struct {
	Name    string
	Methods []string
	Path    string
	Handler http.HandlerFunc
}

var routes = []Route{
	{
		Name:    "Health OK",
		Methods: []string{http.MethodGet},
		Path:    "/ok",
		Handler: handlers.HandleAdapter(handlers.HandleOK),
	}, {
		Name:    "List Continents",
		Methods: []string{http.MethodGet},
		Path:    "/continents",
		Handler: handlers.HandleAdapter(handlers.HandleListContinents),
	}, {
		Name:    "List Countries",
		Methods: []string{http.MethodGet},
		Path:    "/countries",
		Handler: handlers.HandleAdapter(handlers.HandleListCountries),
	}, {
		Name:    "Get Country Info",
		Methods: []string{http.MethodGet},
		Path:    "/countries/{code}",
		Handler: handlers.HandleAdapter(handlers.HandleGetCountryInfo),
	},
}
