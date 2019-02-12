package main

import (
	"flag"
	"log"

	"./migration"
	"./routes"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la base de datos")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzó la migración...")
		migration.Migrate()
		log.Println("Finalizó la migración...")
	}

	// Init routes
	routes.InitRoutes()
}
