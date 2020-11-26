package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

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
