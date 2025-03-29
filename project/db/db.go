package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("DB connection failed")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
     CREATE TABLE IF NOT EXISTS events(
	  id INTEGER PRIMARY KEY AUTOINCREMENT,
	  name TEXT NOT NULL,
	  email TEXT NOT NULL,
	  location TEXT NOT NULL,
	  date DATETIME NOT NULL,
	  user_id INTEGER
	  )
	`
	_, err := DB.Exec(createEventsTable)

	if err != nil {
		log.Println(err)
		panic("Could not create events")
	}
}
