package csv

import (
	"encoding/csv"
	"github.com/bhavye-omniful/GORM/models"
	"github.com/gocarina/gocsv"
	"log"
	"os"
)

func Init() {
	CreateCsvFile("example.csv")
}

func CreateCsvFile(name string) {
	csvFile, err := os.Create(name)
	if err != nil {
		log.Fatal("Error creating a csv file")
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Error closing csv file")
		}
	}(csvFile)
}

//func WriteHeadersToCSV(header []string) error {
//	file, err := os.Open(CsvObj.File.Name())
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//
//	err = CsvObj.Writer.Write(header)
//	if err != nil {
//		log.Fatal("Error writing header to CSV:", err)
//		return err
//	}
//	return nil
//}

//func WriteToCSV(data []string) error {
//	file, err := os.Open(CsvObj.File.Name())
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//
//	err = CsvObj.Writer.Write(data)
//	if err != nil {
//		log.Fatal("Error writing data to CSV:", err)
//		return err
//	}
//	return nil
//}

func WriteAllToCsv(data [][]string) error {
	file, err := os.OpenFile("example.csv", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(data)
	if err != nil {
		log.Fatal("Error writing data to CSV:", err)
		return err
	}
	return nil
}

func ReadFromCSV(csvFile string) (data []*models.Source, err error) {
	file, err := os.Open(csvFile)
	if err != nil {
		log.Fatal("Error opening CSV file:", err)
		return nil, err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal("Error closing CSV file:", err)
		}
	}(file)

	err = gocsv.UnmarshalFile(file, &data)
	if err != nil {
		log.Fatal("Error unmarshalling file")
		return nil, err
	}
	return
}
