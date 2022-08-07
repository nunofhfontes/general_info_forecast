package datamodels

import "time"

type Account struct {
	Id       int64
	Name     string    `xorm:"varchar(25) not null unique 'name' "`
	Role     string    `xorm:"varchar(25) not null unique 'role' "`
	Password string    `xorm:"varchar(200)"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}
