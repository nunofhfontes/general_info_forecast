package datamodels

import "time"

type Portfolio struct {
	Id       int64
	UserId   int64
	Type     string    // needed?
	Security string    // this should be a multi valur column, json, jsonb, array
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}
