package main

import (
	"github.com/bhavye-omniful/GORM/database"
	"github.com/bhavye-omniful/GORM/redis"
	"github.com/bhavye-omniful/GORM/routes"
	"github.com/bhavye-omniful/GORM/sqs"
)

func main() {
	redis.Init()
	database.Init()
	sqs.Init()
	routes.Routes()
}
