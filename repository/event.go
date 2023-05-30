package repository

import (
	"database/sql"

	"github.com/maiconpavi/blankfactor-go-test/schema"
	_ "github.com/mattn/go-sqlite3"
)

const (
	// default database file
	defaultFile string = "events.db"
	// creates the events table
	createQuery string = `
  CREATE TABLE IF NOT EXISTS events (
  id INTEGER NOT NULL PRIMARY KEY,
  title TEXT,
  start_time DATETIME NOT NULL,
  end_time DATETIME NOT NULL
  );`
	// list all events that overlap
	listOverlapPairsQuery string = `
	 SELECT e1.id AS id1,
			e1.title AS title1,
			e1.start_time AS start_time1,
			e1.end_time AS end_time1,
			e2.id AS id2,
			e2.title AS title2,
			e2.start_time AS start_time2,
			e2.end_time AS end_time2
		FROM EVENTS e1
		INNER JOIN EVENTS e2 ON e2.end_time >= e1.start_time
		AND e1.end_time >= e2.start_time
		AND e1.id < e2.id`
	// list all events
	listQuery string = `
	SELECT * FROM events
	`
	// get a event by id
	getQuery string = `
	SELECT * FROM events WHERE id = ?
	`
	// delete a event by id
	deleteQuery string = `
	DELETE FROM events WHERE id = ?
	`
	// update a event
	updateQuery string = `
	UPDATE events SET title = ?, start_time = ?, end_time = ? WHERE id = ?
	`
)

// Event Repository Interface for CRUD operations
type EventRepository interface {
	// Insert a new event
	Insert(event schema.Event) (int, error)
	// List all events that overlap
	ListOverlapPairs() ([]schema.EventPair, error)
	// List all events
	List() ([]schema.Event, error)
	// Get a event by id
	Get(id int) (schema.Event, error)
	// Delete a event by id
	Delete(id int) error
	// Update a event
	Update(event schema.Event) error
}

type eventRepository struct {
	db *sql.DB
}

// NewEventRepository creates a new EventRepository
func NewEventRepository(files ...string) (EventRepository, error) {
	var file string
	if len(files) > 0 {
		file = files[0]
	} else {
		file = defaultFile
	}
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(createQuery); err != nil {
		return nil, err
	}
	return &eventRepository{
		db: db,
	}, nil
}

// Insert a new event
func (c *eventRepository) Insert(event schema.Event) (int, error) {
	res, err := c.db.Exec("INSERT INTO events VALUES(NULL,?,?,?);", event.Title, event.StartTime, event.EndTime)
	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}

// List all events that overlap
func (c *eventRepository) ListOverlapPairs() ([]schema.EventPair, error) {
	rows, err := c.db.Query(listOverlapPairsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pairs []schema.EventPair
	for rows.Next() {
		var pair schema.EventPair
		if err := rows.Scan(&pair.ID1, &pair.Title1, &pair.StartTime1, &pair.EndTime1, &pair.ID2, &pair.Title2, &pair.StartTime2, &pair.EndTime2); err != nil {
			return nil, err
		}
		pairs = append(pairs, pair)
	}
	return pairs, nil
}

// List all events
func (c *eventRepository) List() ([]schema.Event, error) {
	rows, err := c.db.Query(listQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []schema.Event
	for rows.Next() {
		var event schema.Event
		if err := rows.Scan(&event.ID, &event.Title, &event.StartTime, &event.EndTime); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

// Get a event by id
func (c *eventRepository) Get(id int) (schema.Event, error) {
	var event schema.Event
	err := c.db.QueryRow(getQuery, id).Scan(&event.ID, &event.Title, &event.StartTime, &event.EndTime)
	if err != nil {
		return event, err
	}
	return event, nil
}

// Delete a event by id
func (c *eventRepository) Delete(id int) error {
	_, err := c.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	return nil
}

// Update a event
func (c *eventRepository) Update(event schema.Event) error {
	_, err := c.db.Exec(updateQuery, event.Title, event.StartTime, event.EndTime, event.ID)
	if err != nil {
		return err
	}
	return nil
}

// Close the database connection
func (c *eventRepository) Close() error {
	return c.db.Close()
}
