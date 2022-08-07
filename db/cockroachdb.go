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

}
