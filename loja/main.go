package main

import (
	"GitHub/loja/routes"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
