package models

import (
	"log"
	"time"

	"cloudterms.net/project/db"
)

type Event struct {
	ID       int64
	Name     string    `binding:"required"`
	Email    string    `binding:"required,email"`
	Location string    `binding:"required"`
	Date     time.Time `binding:"required"`
	UserID   int64
}

var events = []Event{}

func (e Event) Save() error {
	query := `INSERT INTO events(name, email, location, date, user_id) 
	VALUES (?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Email, e.Location, e.Date, e.UserID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Email, &event.Location, &event.Date, &event.UserID)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	if err := row.Scan(&event.ID, &event.Name, &event.Email, &event.Location, &event.Date, &event.UserID); err != nil {
		return nil, err
	}

	return &event, nil
}
