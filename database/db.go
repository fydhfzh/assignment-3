package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)


var (
	db *sql.DB
	err error
)

const (
	host = "localhost"
	user = "admin"
	password = "postgres"
	dialect = "postgres"
	port = 5432
	dbname = "weather"
)

func handleDBConnection(){
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, user, password, port, dbname)
	db, err = sql.Open(dialect, psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

}

func createRequiredTables(){
	createWeatherTableQuery := `
		CREATE TABLE IF NOT EXISTS weather (
			id SERIAL NOT NULL,
			wind INT NOT NULL,
			water INT NOT NULL,
			water_status VARCHAR(16) NOT NULL,
			wind_status VARCHAR(16) NOT NULL,
			created_at timestamptz NOT NULL default now(),
			updated_at timestamptz NOT NULL default now()
		);
	`

	_, err := db.Exec(createWeatherTableQuery)

	if err != nil {
		log.Fatal(err)
	}
}

func InitializeDB() {
	handleDBConnection()
	createRequiredTables()
}

func GetInstance() *sql.DB {
	return db
}