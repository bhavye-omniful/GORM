package main

import (
	"github.com/bhavye-omniful/GORM/database"
	"github.com/bhavye-omniful/GORM/routes"
)

func main() {
	database.Init()
	routes.Routes()
}
