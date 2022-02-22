package datamodels

// Locations struct which contains
// an array of locations
type Locations struct {
	Locations []Location `json:"locations"`
}

// Location struct which contains a name
// coordinates, country code/name and an array of people
type Location struct {
	Name    string      `json:"name"`
	Coords  Coordinates `json:"coords"`
	Code    string      `json:"code"`
	Country string      `json:"country"`
	People  []string    `json:"people"`
}

type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}
