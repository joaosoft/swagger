package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"swagger/controllers"
)

var (
	Router = mux.NewRouter()
)

func init() {
	// swagger:route GET /persons/{id_person} person-tag GetPersonByIDRequestParameter
	//
	// Get person by id.
	//
	// This will return the person information.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Deprecated: false
	//
	//     Responses:
	//       default: PersonResponse
	//       200: PersonResponse
	//       400: ErrorResponse
	Router.HandleFunc("/v1/persons/{id_person}", controllers.GetPersonByID).Methods(http.MethodGet)

	// swagger:route GET /persons/{id_person}/addresses/{id_address} address-tag GetPersonAddressByIDRequestParameter
	//
	// Get the person address.
	//
	// This will return error.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Deprecated: false
	//
	//     Responses:
	//       default: AddressResponse
	//       200: AddressResponse
	//       400: ErrorResponse
	Router.HandleFunc("/v1/persons/{id}/addresses/{id_address}", controllers.GetPersonAddressByID).Methods(http.MethodGet)

	// swagger:route GET / not-found-tag
	//
	// Not Found
	//
	// This will return error.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Deprecated: false
	//
	//     Responses:
	//       default: ErrorResponse
	//       404: ErrorResponse
	Router.HandleFunc("/v1/", controllers.NotFound)

	fs := http.FileServer(http.Dir("./spec"))
	Router.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", fs))
}
