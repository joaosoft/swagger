package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GetPersonByIDRequest example
// @Summary Get person by id.
// @Description Get person by id.
// @ID get-person-by-id
// @Accept  json
// @Produce  json
// @Param   id_person      path   string     true  "Person ID"
// @Param   age      query   int     false  "Age"
// @Success 200 {string} PersonResponse	"ok"
// @Failure 400 {object} ErrorResponse "error"
// @Router /persons/{id_person} [get]
func GetPersonByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	age, _ := strconv.Atoi(req.URL.Query().Get("age"))
	request := GetPersonByIDRequest{
		IdPerson: vars["id_person"],
		Age:      age,
	}

	fmt.Printf("> executing get person for id_person: %s", request.IdPerson)

	// ...

	bytes, _ := json.Marshal(
		PersonResponse{
			Id:   request.IdPerson,
			Name: "João Ribeiro",
			Age:  request.Age,
		},
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
