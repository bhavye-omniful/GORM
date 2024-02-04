package main

import (
	"github.com/bhavye-omniful/GORM/csv"
	"log"
)

type Order struct {
	order_number string `csv:"Order Number"`
}

func main() {
	//redis.Init()
	//database.Init()
	//sqs.Init()
	//routes.Routes()
	csv.Init()

	data, err := csv.ReadFromCSV("orders.csv")
	if err != nil {
		log.Fatal("Error reading from csv file")
	}

	err = csv.WriteAllToCsv(data)
	if err != nil {
		log.Fatal("Error writing to csv file")
	}
}
