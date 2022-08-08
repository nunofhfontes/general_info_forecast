package db

import (
	fiberConfig "dataForecast/utils/configs"
	"fmt"
	"log"
	"strconv"
	"xorm.io/xorm"

	_ "github.com/lib/pq"
)

func InitCockroachDB(config fiberConfig.Config) {

	// connect to online cockroach DB -> trader-factory

	fmt.Printf("config.DBDriver -> %s\n", config.DBDriver)
	fmt.Printf("config.DBSource -> %s\n", config.DBSource)

	//engine, err := xorm.NewEngine(driverName, dataSourceName)
	engine, err := xorm.NewEngine(config.DBDriver, config.DBSource)
	engine.ShowSQL(true)

	if err != nil {
		log.Println("engine creation failed", err)
	}

	// checking health
	err = engine.Ping()
	if err != nil {
		panic(err)
	}

	//var now time.Time
	//engine.Raw("SELECT NOW()").Scan(&now)

	// temporary - check if the user table has data
	counts, _ := engine.Count()

	fmt.Printf("count: " + strconv.FormatInt(counts, 10) + "\n")

	// try to insert a user
	fmt.Printf("\ntrying to insert a user \n\n")

	type User struct {
		Id       int64
		Username string
		//Name     string
		//Salt     string
		//Age      int
		Password string `xorm:"varchar(200)"`
		//AccountId int64
		//Created  time.Time `xorm:"created"`
		//Updated  time.Time `xorm:"updated"`
	}

	var existingUser User

	has, err := engine.Get(&existingUser)
	// SELECT * FROM user LIMIT 1
	fmt.Printf("\n Has user?: %t \n\n", has)

	fmt.Printf("\n Has user? existing user ->   %s \n\n", existingUser.Username)

	has1, err := engine.Where("username = ?", "nunofhfontes").Desc("id").Get(&existingUser)
	// SELECT * FROM user WHERE name = ? ORDER BY id DESC LIMIT 1
	fmt.Printf("\n Has user1?: %t \n\n", has1)

	fmt.Printf("\n Has user1? existing user ->   %s \n\n", existingUser.Password)

	user := User{Username: "xiaoxiao", Password: "lunny"} // , Created: time.Now()
	affected, err := engine.Insert(&user)

	fmt.Printf("\n Affected rows: %d \n\n", affected)
}
