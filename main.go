package main

import (
	"net/http"

	"github.com/goweb/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
