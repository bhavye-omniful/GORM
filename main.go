package main

import (
	"github.com/bhavye-omniful/GORM/csv"
	"github.com/gocarina/gocsv"
	"os"
)

type Order struct {
	OrderNumber string `csv:"Order Number"`
	Sku         string `csv:"SKU"`
	Comment     string `csv:"BlaBla"`
}

func main() {
	//redis.Init()
	//database.Init()
	//sqs.Init()
	//routes.Routes()
	csv.Init()

	//data, err := csv.ReadFromCSV("orders.csv")
	//if err != nil {
	//	log.Fatal("Error reading from csv file")
	//}

	data := []*Order{&Order{
		OrderNumber: "1234",
		Sku:         "9999",
		Comment:     "testing",
	}}

	//err = csv.WriteAllToCsv(data)
	//if err != nil {
	//	log.Fatal("Error writing to csv file")
	//}

	file, err := os.OpenFile("example.csv", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	err = gocsv.MarshalFile(data, file)
	if err != nil {
		return
	}

}
