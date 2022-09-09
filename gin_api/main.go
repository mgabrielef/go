package main

import (
	"github.com/mgabrielef/GOLANG/database"
	"github.com/mgabrielef/GOLANG/routes"
)

func main() {
	database.DbConnection()
	routes.HandleRequests()
}
