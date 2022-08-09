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
		Id    int64
		Email string
		//Username string
		Name string
		//FirstName int64 `xorm:"first_name"`
		//LastName  int64 `xorm:"last_name"`
		//Name     string
		//Salt     string
		//Age      int
		Password string `xorm:"varchar(200)"`
		//AccountId int64  `xorm:"account_id"`
		//AccountId int64
		//Created  time.Time `xorm:"created"`
		//Updated  time.Time `xorm:"updated"`
	}

	type Test struct {
		Id    int64
		Descr string
	}

	var existingUser User

	has, err := engine.Get(&existingUser)
	// SELECT * FROM user LIMIT 1public
	fmt.Printf("\n Has user?: %t \n\n", has)

	fmt.Printf("\n Has user? existing user ->   %s \n\n", existingUser.Name)

	has1, err := engine.Where("username = ?", "nunofhfontes").Desc("id").Get(&existingUser)
	// SELECT * FROM user WHERE name = ? ORDER BY id DESC LIMIT 1
	fmt.Printf("\n Has user1?: %t \n\n", has1)

	fmt.Printf("\n Has user1? existing user ->   %s \n\n", existingUser.Password)

	user := User{
		Name:     "xiaoxiao",
		Email:    "nunof.h.fontes@gmail.com",
		Password: "lunny",
		//AccountId: 1} // , Created: time.Now()
	}

	affected, err := engine.Insert(&user)

	fmt.Printf("\n Affected rows: %d \n\n", affected)

	testVar := Test{
		Id:    110,
		Descr: "a new variable - override",
	}
	affectedTest, err := engine.Insert(&testVar)

	fmt.Printf("\n Affected test rows: %d \n\n", affectedTest)
}
