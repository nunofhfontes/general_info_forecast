package fileUtils

import (
	"fmt"
	"os"
)

func ReadLocations() {
	println("reading json locations files")

	// Open our jsonFile
	jsonFile, err := os.Open("./data/locations.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened locations.json")

	// print file content
	println(jsonFile.)

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

}
