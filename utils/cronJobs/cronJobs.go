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

		// TODO - add coordenates
		resp, err := http.Get("https://www.7timer.info/bin/api.pl?lon=-9.29&lat=39.25&product=two&output=json")
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
