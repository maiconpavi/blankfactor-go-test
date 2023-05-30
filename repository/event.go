package repository

import (
	"database/sql"

	"github.com/maiconpavi/blankfactor-go-test/schema"
	_ "github.com/mattn/go-sqlite3"
)

const (
	defaultFile string = "events.db"
	createQuery string = `
  CREATE TABLE IF NOT EXISTS events (
  id INTEGER NOT NULL PRIMARY KEY,
  title TEXT,
  start_time DATETIME NOT NULL,
  end_time DATETIME NOT NULL
  );`
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
		AND e1.id != e2.id
		AND e1.start_time < e2.start_time`
	listQuery string = `
	SELECT * FROM events
	`
	getQuery string = `
	SELECT * FROM events WHERE id = ?
	`
	deleteQuery string = `
	DELETE FROM events WHERE id = ?
	`
	updateQuery string = `
	UPDATE events SET title = ?, start_time = ?, end_time = ? WHERE id = ?
	`
)

type EventRepository interface {
	Insert(event schema.Event) (int, error)
	ListOverlapPairs() ([]schema.EventPair, error)
	List() ([]schema.Event, error)
	Get(id int) (schema.Event, error)
	Delete(id int) error
	Update(event schema.Event) error
}

type eventRepository struct {
	db *sql.DB
}

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

func (c *eventRepository) Get(id int) (schema.Event, error) {
	var event schema.Event
	err := c.db.QueryRow(getQuery, id).Scan(&event.ID, &event.Title, &event.StartTime, &event.EndTime)
	if err != nil {
		return event, err
	}
	return event, nil
}

func (c *eventRepository) Delete(id int) error {
	_, err := c.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	return nil
}

func (c *eventRepository) Update(event schema.Event) error {
	_, err := c.db.Exec(updateQuery, event.Title, event.StartTime, event.EndTime, event.ID)
	if err != nil {
		return err
	}
	return nil
}

func (c *eventRepository) Close() error {
	return c.db.Close()
}
