package main

import (
	fUtils "dataForecast/utils/files"
	"fmt"
)

func main() {
	fmt.Println("A cron job")

	// 1 - read json with locations converting to a raw variable
	fUtils.ReadLocationsRaw()

	// 2 - read json with locations parsing info to struts
	fUtils.ReadLocationsParsed4Strut()

	// 2 - init cron jobs
	//cronJobs.InitCronJobs()
}
