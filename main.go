package main

import (
	"github.com/PedroVMB/API-GO-CHECKLIST-CONDOMINIO/database"
	"github.com/PedroVMB/API-GO-CHECKLIST-CONDOMINIO/routes"
)

func main() {
	database.Connection()
	routes.HandleRequests()
}