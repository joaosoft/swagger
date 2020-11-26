package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"swagger/go-swagger/controllers"
)

var (
	Router = mux.NewRouter()
)

func init() {

	Router.HandleFunc("/v1/persons/{id_person}", controllers.GetPersonByID).Methods(http.MethodGet)
	Router.HandleFunc("/v1/persons/{id}/addresses/{id_address}", controllers.GetPersonAddressByID).Methods(http.MethodGet)
	Router.HandleFunc("/v1/errors", controllers.GetErrorByID)

	fs := http.FileServer(http.Dir("./spec"))
	Router.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", fs))
}
