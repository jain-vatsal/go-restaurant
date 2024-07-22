package main

import (
	"os"

	"github.com/jain-vatsal/go-restaurant/constants"
	"github.com/jain-vatsal/go-restaurant/routes"
)

func main() {

	port := os.Getenv(constants.PORT)
	if port == "" {
		port = constants.DEFAULT_PORT
	}

	routes.SetUpRouter(port)
}
