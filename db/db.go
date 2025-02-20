package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type Todo struct {
	ID          string `json:"id,omitempty"`
	Task        string `json:"task"`
	Author      string `json:"author"`
	CreatedDate string `json:"created_date,omitempty"`
	UpdateDate  string `json:"updated_date,omitempty"`
	Completed   bool   `json:"completed,omitempty"`
}

func getConnectionString(create bool) string {
	var (
		host = os.Getenv("DB_HOST")
		port = os.Getenv("DB_PORT")
		dbUser = os.Getenv("DB_USER")
		dbName = os.Getenv("DB_NAME")
		dbPassword = os.Getenv("DB_PASSWORD")
	)

	connection := ""
	if create {
		connection = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s sslmode=disable",
			host,
			port,
			dbUser,
			dbPassword,
		)
	} else {
		connection = fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			host,
			port,
			dbUser,
			dbName,
			dbPassword,
		)
	}

	return connection
}

func handleDbConnectionError(err error) {
 	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func Init() {
	dbConn := getConnectionString(false)

	db, err = gorm.Open("postgres", dbConn)

	if err != nil {
		log.Println("Database doesn't exisiting auto-creating ...")
		createDbConn := getConnectionString(true)
		db, err = gorm.Open("postgres", createDbConn)
		if err != nil {
			handleDbConnectionError(err)
		}

		query := fmt.Sprintf("CREATE DATABASE %s;", os.Getenv("DB_NAME"))
		db = db.Exec(query)
		db.AutoMigrate(Todo{})
	}

	handleDbConnectionError(err)

	log.Println("Database exisit ...")
	db.AutoMigrate(Todo{})
}