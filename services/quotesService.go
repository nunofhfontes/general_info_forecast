package services

import (
	"context"
	model "dataForecast/models"
	"fmt"
	"github.com/carlmjohnson/requests"
	"io/ioutil"
	"log"
	"net/http"
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

func TestGetStockPriceFinModPre() {
	url := "https://financialmodelingprep.com/api/v3/quote/AAPL,FB?apikey=a12a1c9244f3a11dfdb9de8a3b5f7bdc"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)

	fmt.Println("===> " + responseString)
}
