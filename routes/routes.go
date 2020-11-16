// Package Example API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost:8090
//     BasePath: /v1
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: John Doe<john.doe@example.com> http://john.doe.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Extensions:
//     x-meta-value: value
//     x-meta-array:
//       - value1
//       - value2
//     x-meta-array-obj:
//       - name: obj
//         value: field
//
// swagger:meta
package routes

import (
	"net/http"
	"swagger/controllers"
)

func init() {
	// swagger:route GET /success success
	//
	// Get success.
	//
	// This will return success.
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
	//       default: SuccessResponse
	//       200: SuccessResponse
	//       400: ErrorResponse
	http.HandleFunc("/v1/success", controllers.HandleSuccess)

	// swagger:route GET /error error
	//
	// Get error.
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
	//       200: ErrorResponse
	//       400: ErrorResponse
	http.HandleFunc("/v1/error", controllers.HandleError)

	fs := http.FileServer(http.Dir("./spec"))
	http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))
}
