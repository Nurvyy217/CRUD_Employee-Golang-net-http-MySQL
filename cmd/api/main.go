package main

import (

	"net/http"

	"github.com/Nurvyy217/crud-employee-go/external/database"
	"github.com/Nurvyy217/crud-employee-go/routes"
)


func main() {
	db := database.InitDatabase()
	server := http.NewServeMux()
	routes.MapRoutes(server, db)
	http.ListenAndServe(":8080", server)
}