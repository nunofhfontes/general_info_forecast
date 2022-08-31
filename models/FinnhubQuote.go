package datamodels

import "time"

type FinnhubQuote struct {
	Timestamp          int32   `json:"t"`
	CurrentPrice       float32 `xorm:"varchar(25) not null unique 'name'" json:"c"`
	Change             float32 `xorm:"varchar(25) not null unique 'ticker'" json:"d"`
	PercentChange      float32 `xorm:"INT not null unique 'current_value'" json:"dp"`
	HighPriceOfTheDay  float32 `json:"h"`
	LowPriceOfTheDay   float32 `json:"l"`
	OpenPriceOfTheDay  float32 `json:"o"`
	PreviousClosePrice float32 `json:"opc"`

	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`

	//TODO
	// HOW TO TRACK Security value, how to keep historical price??
}

//c
//Current price
//
//d
//Change
//
//dp
//Percent change
//
//h
//High price of the day
//
//l
//Low price of the day
//
//o
//Open price of the day
//
//pc
//Previous close price
