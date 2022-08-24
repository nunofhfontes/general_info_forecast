package services

import (
	"context"
	model "dataForecast/models"
	"fmt"
	"github.com/carlmjohnson/requests"
	"log"
)

func GetQuote() {

	fmt.Printf("Fetching data \n")
	var financialData model.Security
	err := requests.
		URL("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=IBM&interval=5min&apikey=demo").
		ToJSON(&financialData).
		Fetch(context.Background())

	if err != nil {
		log.Println("http client failed", err)
	}

	log.Println("After Data Fetch")

	log.Println("result: ", financialData.Name)
}
