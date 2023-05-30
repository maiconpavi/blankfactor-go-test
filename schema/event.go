package schema

import (
	"time"
)

type Event struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

func NewEvent(title string, startTime, endTime time.Time) Event {
	return Event{
		ID:        0,
		Title:     title,
		StartTime: startTime,
		EndTime:   endTime,
	}
}

type EventPair struct {
	ID1        int       `json:"id1"`
	Title1     string    `json:"title1"`
	StartTime1 time.Time `json:"start_time1"`
	EndTime1   time.Time `json:"end_time1"`
	ID2        int       `json:"id2"`
	Title2     string    `json:"title2"`
	StartTime2 time.Time `json:"start_time2"`
	EndTime2   time.Time `json:"end_time2"`
}
