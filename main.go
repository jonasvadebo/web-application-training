package main

import (
	"net/http"

	controller "github.com/jonasvadebo/web-application-training/controllers"
)

func main() {
	controller.RegisterControllers()
	http.ListenAndServe(":8080", nil)
}
