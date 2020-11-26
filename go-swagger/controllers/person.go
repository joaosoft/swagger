package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

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
func GetPersonByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	fmt.Printf("> executing get person for id_person: %s", vars["id_person"])

	age, _ := strconv.Atoi(req.URL.Query().Get("age"))
	bytes, _ := json.Marshal(
		PersonResponse{
			Id:   vars["id_person"],
			Name: "João Ribeiro",
			Age:  age,
		},
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
