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

	// swagger:operation GET /errors getErrors
	//
	// Returns the error definition
	//
	// Could be any error
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: id_error
	//   in: query
	//   description: the error id
	//   required: true
	//   type: integer
	//   format: int32
	// responses:
	//   '200':
	//     description: the error definition
	//     schema:
	//       "$ref": "#/responses/ErrorResponse"
	//   '204':
	//     description: no content
	//   default:
	//     description: the error definition
	//     schema:
	//       "$ref": "#/responses/ErrorResponse"
	Router.HandleFunc("/v1/errors", controllers.GetErrorByID)

	fs := http.FileServer(http.Dir("./spec"))
	Router.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", fs))
}
