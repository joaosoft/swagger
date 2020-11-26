package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// GetErrorByID example
// @Summary Returns the error definition
// @Description Returns the error definition
// @ID get-error-by-id
// @Accept  json
// @Produce  json
// @Param   id_error      query   int     true  "Error ID"
// @Success 200 {object} ErrorResponse "ok"
// @Success 204 {string} string	"ok"
// @Router /errors [get]
func GetErrorByID(w http.ResponseWriter, req *http.Request) {
	errorID, _ := strconv.Atoi(req.URL.Query().Get("id_error"))
	fmt.Printf("> executing get errors for id: %d", errorID)

	statusText := http.StatusText(errorID)

	if statusText != "" {
		bytes, _ := json.Marshal(
			ErrorResponse{
				Code:    errorID,
				Message: statusText,
			},
		)
		w.Write(bytes)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}

	w.Header().Set("Content-Type", "application/json")
}
