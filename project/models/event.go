package models

import "time"

type Event struct {
	ID       int
	Name     string    `binding:"required"`
	Email    string    `binding:"required,email"`
	Location string    `binding:"required"`
	Date     time.Time `binding:"required"`
	UserID   int
}

var events = []Event{}

func (e Event) Save() {
	// later: add database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
