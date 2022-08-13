package services

import (
	"context"
	model "dataForecast/models"
	"github.com/carlmjohnson/requests"
)

func getQuote() {

	var finacialData model.Security
	err := requests.
		URL("https://example.com/my-json").
		ToJSON(&finacialData).
		Fetch(context.Background())
}
