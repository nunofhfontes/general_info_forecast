package misc

import (
	"github.com/go-co-op/gocron"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func InitCronJobs() {
	println("Initiating cron jobs...")

	s := gocron.NewScheduler(time.UTC)

	s.Every(30).Seconds().Do(func() {
		println("every 10 seconds print cron")

		resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=lourinha&appid=62d2c58c3bf1cda8c560ff1a7606e232&units=metric")
		if err != nil {
			log.Fatalln(err)
		}

		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		//Convert the body to type string
		sb := string(body)
		log.Printf(sb)
	})

	// start the cron blocking the current execution
	s.StartBlocking()

	// cron expressions supported
	//s.Cron("*/1 * * * *").Do(func() {
	//	println("every minute")
	//})
	//})
}
