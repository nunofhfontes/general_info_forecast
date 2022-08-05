package db

import (
	fiberConfig "dataForecast/utils/configs"
	"log"
	"xorm.io/xorm"
)

func InitCockroachDB(config fiberConfig.Config) {

	// connect to online cockroach DB -> trader-factory

	//engine, err := xorm.NewEngine(driverName, dataSourceName)
	engine, err := xorm.NewEngine(config.DBDriver, config.DBSource)

	if err != nil {
		log.Println("engine creation failed", err)
	}

	// checking health
	err = engine.Ping()
	if err != nil {
		panic(err)
	}

}
