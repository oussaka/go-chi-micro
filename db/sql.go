package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/oussaka/go-chi-micro/model"
	log "github.com/sirupsen/logrus"
	"os"
)

type SqlCon struct {
	DbPool *gorm.DB
}

var connector *SqlCon

//type sqlConn struct {
//	DbPool *sql.DB
//}

//var connector *sqlConn

var counts int64

func InitPgsql() *SqlCon {
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
func initDB() (*SqlCon, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), "go-db", os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.Blogs{})

	return &SqlCon{db}, nil
}

/* func CreateConnection() (*sqlConn, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", "postgres", "postgres", "pgsql", "postgres")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if connector != nil {
		log.Info("DataBase is initialized")
		return connector, nil
	}
	log.Info("DataBase was not initialized ..initializing again")

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	connector, err = &sqlConn{db}, err

	return connector, nil
} */

func GetDBConnection() *gorm.DB {
	return connector.DbPool
}
