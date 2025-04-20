package service

import (
	"fmt"
	"gorm.io/gorm"
)

var connector *sqlConn

func InitMysql() *sqlConn {
	if connector != nil {
		log.Info("DataBase is initialized")
		return connector
	}
	log.Info("DataBase was not initialized ..initializing again")
	var err error
	connector, err = initDB()
	if err != nil {
		panic(err)
	}
	return connector
}

// DB Initialization

func initDB() (*sqlConn, error) {
	log.Info("connection details => ", config.GetYamlValues().DBConfig, config.GetYamlValues().DBConfig.Port)
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		config.GetYamlValues().DBConfig.Server, config.GetYamlValues().DBConfig.Username, config.GetYamlValues().DBConfig.Schema, config.GetYamlValues().DBConfig.Password) //Build connection string

	db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		panic(err)
	}
	if maxCons := config.GetYamlValues().DBConfig.MaxConnection; maxCons > 0 {
		db.DB().SetMaxOpenConns(maxCons)
		db.DB().SetMaxIdleConns(maxCons / 3)
	}

	return &sqlConn{db}, nil
}

func NewClient(config *Config) SqlClient {
	return &sqlClient{
		config: config,
	}
}
