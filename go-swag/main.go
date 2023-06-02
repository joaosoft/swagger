package main

// @title Package Person API.
// @version 0.0.1
// @description The purpose of this application is to provide an application that is using plain go code to define an API
// @termsOfService http://swagger.io/terms/

// @contact.name Contact
// @contact.url https://github.com/joaosoft
// @contact.email joaosoft@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081
// @BasePath /v1

import (
	"fmt"
	"github.com/joaosoft/swagger/go-swag/routes"
	_ "github.com/joaosoft/swagger/go-swag/routes"
	"net/http"
)

func main() {
	fmt.Println("server started at http://localhost:8081/swaggerui/")
	if err := http.ListenAndServe(":8081", routes.Router); err != nil {
		panic(err)
	}
}
