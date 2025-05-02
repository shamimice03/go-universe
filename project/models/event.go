package models

import (
	"log"
	"time"

	"cloudterms.net/project/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Date        time.Time `binding:"required"`
	UserID      int64
}

var events = []Event{}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, date, user_id) 
	VALUES (?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.Date, e.UserID)
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

func (e *Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, date = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	stmt.Exec(e.Name, e.Description, e.Location, e.Date, e.ID)

	return err
}

func (e *Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
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
	events = []Event{}

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserID)

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
	if err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserID); err != nil {
		return nil, err
	}

	return &event, nil
}

func (e *Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e *Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)

	return err
}
