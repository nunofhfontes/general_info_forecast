package fileUtils

import (
	datamodels "dataForecast/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadLocationsRaw() {
	println("reading json locations files")

	// Open our jsonFile
	jsonFile, err := os.Open("./data/locations.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened locations.json")

	// print file content
	println(jsonFile)

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// declare an interface
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result["locations"])
}

func ReadLocationsParsed4Strut() {
	println("reading json locations files")

	// Open our jsonFile
	jsonFile, err := os.Open("./data/locations.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened locations.json")

	// print file content
	println(jsonFile)

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we declare our Locations array
	var locations datamodels.Locations

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'locations' which we defined above
	json.Unmarshal(byteValue, &locations)

	// we iterate through every location within our locations array and
	// print out the locations details
	for i := 0; i < len(locations.Locations); i++ {
		fmt.Println("User Type: " + locations.Locations[i].Name)
	}

	mapOfLocationAndPeople := make(map[string][]string)
	println(mapOfLocationAndPeople)

	mapOfLocationAndWeatherStatement := make(map[string][]string)
	println(mapOfLocationAndWeatherStatement)
}

// TODO - call this
// 1 - read json with locations converting to a raw variable
//fUtils.ReadLocationsRaw()

// 2 - read json with locations parsing info to struts
//fUtils.ReadLocationsParsed4Strut()

// 2 - init cron jobs
//cronJobs.InitCronJobs()
