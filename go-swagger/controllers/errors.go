package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

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
