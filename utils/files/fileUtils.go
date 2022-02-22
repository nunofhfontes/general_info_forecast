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
		//fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		//fmt.Println("User Name: " + users.Users[i].Name)
		//fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}
}
