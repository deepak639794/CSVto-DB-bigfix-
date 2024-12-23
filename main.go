package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CSV_data struct {
	SiteID                string
	FixletID              string
	Name                  string
	Criticality           string
	RelevantComputerCount int
}

func main() {
	dbcon := "host=postgres user=deepakAgrawl password=dee@123 dbname=CSV_db port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dbcon), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&CSV_data{}) //

	file, err := os.Open("../../Files/fixlets.csv")
	if err != nil {
		log.Fatalf("Not able to read CSV file %v", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV file data %v", err)
	}

	for _, i := range records[1:] { // Skip header
		r := CSV_data{
			SiteID:                i[0],
			FixletID:              i[1],
			Name:                  i[2],
			Criticality:           i[3],
			RelevantComputerCount: atoi(i[4]),
		}
		db.Create(&r)
	}

	fmt.Println("CSV data imported successfully.")
}
func atoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
