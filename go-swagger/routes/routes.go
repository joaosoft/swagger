package routes

import (
	"github.com/gorilla/mux"
	"github.com/joaosoft/swagger/go-swagger/controllers"
	"net/http"
)

var (
	Router = mux.NewRouter()
)

func init() {

	Router.HandleFunc("/v1/persons/{id_person}", controllers.GetPersonByID).Methods(http.MethodGet)
	Router.HandleFunc("/v1/persons/{id_person}/addresses/{id_address}", controllers.GetPersonAddressByID).Methods(http.MethodGet)
	Router.HandleFunc("/v1/errors", controllers.GetErrorByID)

	fs := http.FileServer(http.Dir("./spec"))
	Router.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", fs))
}
