package main

import (
	"fmt"

	"github.com/lucasstettner/api-boilerplate/app"
)

// @title Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	a := app.App{}

	a.Start(true)

	fmt.Println("server has started")
}
