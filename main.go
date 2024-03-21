package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bhavye-omniful/GORM/csv"
	"github.com/bhavye-omniful/GORM/database"
	"github.com/bhavye-omniful/GORM/redis"
	"github.com/bhavye-omniful/GORM/routes"
	"github.com/bhavye-omniful/GORM/sqs"
	"github.com/gocarina/gocsv"
	"os"
)

type Order struct {
	OrderNumber string `csv:"Order Number"`
	Sku         string `csv:"SKU"`
	Comment     string `csv:"BlaBla"`
}

type MyEventSlice []MyEvent
type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event *MyEvent) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}
	message := fmt.Sprintf("Hello %s!", event.Name)
	return &message, nil
}

func main() {
	redis.Init()
	database.Init()
	sqs.Init()
	routes.Routes()
	csv.Init()

	lambda.Start(HandleRequest)

	//data, err := csv.ReadFromCSV("orders.csv")
	//if err != nil {
	//	log.Fatal("Error reading from csv file")
	//}

	data := []*Order{&Order{
		OrderNumber: "1234",
		Sku:         "123143",
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
