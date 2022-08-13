package datamodels

import "time"

type Security struct {
	Id           int64
	Name         string    `xorm:"varchar(25) not null unique 'name' "`
	ticker       string    `xorm:"varchar(25) not null unique 'ticker' "`
	CurrentValue float64   `xorm:"INT not null unique 'current_value' "`
	Created      time.Time `xorm:"created"`
	Updated      time.Time `xorm:"updated"`
}
