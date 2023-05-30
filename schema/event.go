package schema

import (
	"time"
)

// Event struct for event
type Event struct {
	// Event id
	ID int `json:"id"`
	// Eent title
	Title string `json:"title"`
	// Event start time
	StartTime time.Time `json:"start_time"`
	// Event end time
	EndTime time.Time `json:"end_time"`
}

// Creates a new Event
func NewEvent(title string, startTime, endTime time.Time) Event {
	return Event{
		ID:        0,
		Title:     title,
		StartTime: startTime,
		EndTime:   endTime,
	}
}

// event pair that overlap
type EventPair struct {
	// First event id
	ID1 int `json:"id1"`
	// First event title
	Title1 string `json:"title1"`
	// First event start time
	StartTime1 time.Time `json:"start_time1"`
	// First event end time
	EndTime1 time.Time `json:"end_time1"`
	// Second event id
	ID2 int `json:"id2"`
	// Second event title
	Title2 string `json:"title2"`
	// Second event start time
	StartTime2 time.Time `json:"start_time2"`
	// Second event end time
	EndTime2 time.Time `json:"end_time2"`
}
