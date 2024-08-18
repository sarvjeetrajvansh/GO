package models

import (
	"time"
)

type Event struct {
	ID          int
	UserId      int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Allocation  string `binding:"required"`
	DateTime    time.Time
}

var events = []Event{}

func (e Event) Save() {
	// TODO: Add it to DB
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
