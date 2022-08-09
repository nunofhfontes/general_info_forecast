package datamodels

type User struct {
	Id       int64
	Username string
	//Name     string
	//Salt     string
	//Age      int
	Password string `xorm:"varchar(200)"`
	//Created  time.Time `xorm:"created"`
	//Updated  time.Time `xorm:"updated"`
}
