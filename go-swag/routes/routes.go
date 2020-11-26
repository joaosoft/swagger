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

	// GetPersonByIDRequestParameter example
	// @Summary Get person by id.
	// @Description Get person by id.
	// @ID get-person-by-id
	// @Accept  json
	// @Produce  json
	// @Param   id_person      path   int     true  "Person ID"
	// @Success 200 {string} PersonResponse	"ok"
	// @Failure 400 {object} ErrorResponse "error"
	// @Router /persons/{id_person} [get]
	Router.HandleFunc("/v1/persons/{id_person}", controllers.GetPersonByID).Methods(http.MethodGet)

	// GetPersonAddressByIDRequestParameter example
	// @Summary Get the person address.
	// @Description Get the person address.
	// @ID get-person-address-by-id
	// @Accept  json
	// @Produce  json
	// @Param   id_person      path   int     true  "Person ID"
	// @Param   id_address      path   int     true  "Person ID"
	// @Success 200 {string} AddressResponse	"ok"
	// @Failure 400 {object} ErrorResponse "error"
	// @Router /persons/{id_person}/addresses/{id_address} [get]
	Router.HandleFunc("/v1/persons/{id}/addresses/{id_address}", controllers.GetPersonAddressByID).Methods(http.MethodGet)

	// GetErrorByID example
	// @Summary Returns the error definition
	// @Description Returns the error definition
	// @ID get-error-by-id
	// @Accept  json
	// @Produce  json
	// @Param   id_error      query   int     true  "Error ID"
	// @Success 200 {object} ErrorResponse "ok"
	// @Failure 204 {string} "no content"
	// @Router /errors [get]
	Router.HandleFunc("/v1/errors", controllers.GetErrorByID)

	fs := http.FileServer(http.Dir("./spec"))
	Router.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", fs))
}
