package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleSuccess(w http.ResponseWriter, req *http.Request) {
	fmt.Println("> executing success")
	bytes, _ := json.Marshal(
		SuccessResponse{
			Success: true,
		},
	)

	fmt.Fprint(w, string(bytes))
}
