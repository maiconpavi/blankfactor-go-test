package repository_test

import (
	"os"
	"testing"
	"time"

	"github.com/maiconpavi/blankfactor-go-test/repository"
	"github.com/maiconpavi/blankfactor-go-test/schema"
)

func getTestRepository() (repository.EventRepository, error) {
	os.Remove("events_test.db")
	return repository.NewEventRepository("events_test.db")
}

func getTestEvent() schema.Event {
	return schema.Event{
		Title:     "Test",
		StartTime: time.Now(),
		EndTime:   time.Now().Add(time.Hour),
	}
}

func TestEventRepositoryInsert(t *testing.T) {
	repo, err := getTestRepository()
	if err != nil {
		t.Error(err)
		return
	}
	event := getTestEvent()
	id, err := repo.Insert(event)
	if err != nil {
		t.Error(err)
		return
	}
	if id == 0 {
		t.Error("id is 0")
		return
	}
}

func TestEventRepositoryList(t *testing.T) {
	repo, err := getTestRepository()
	if err != nil {
		t.Error(err)
		return
	}
	events, err := repo.List()
	if err != nil {
		t.Error(err)
		return
	}
	if len(events) != 0 {
		t.Error("events not is empty")
		return
	}
	event := getTestEvent()

	id, err := repo.Insert(event)
	if err != nil {
		t.Error(err)
		return
	}
	if id == 0 {
		t.Error("id is 0")
		return
	}

	events, err = repo.List()
	if err != nil {
		t.Error(err)
		return
	}
	if len(events) != 1 {
		t.Error("events is not 1")
		return
	}
}

func TestEventRepositoryListOverlapPairs(t *testing.T) {
	event1 := getTestEvent()
	event2 := schema.Event{
		Title:     "Test",
		StartTime: time.Now().Add(time.Minute * 30),
		EndTime:   time.Now().Add(time.Minute * 90),
	}
	event3 := schema.Event{
		Title:     "Test",
		StartTime: time.Now().Add(-time.Minute * 30),
		EndTime:   time.Now().Add(time.Minute * 29),
	}

	repo, err := getTestRepository()
	if err != nil {
		t.Error(err)
		return
	}
	events, err := repo.ListOverlapPairs()
	if err != nil {
		t.Error(err)
		return
	}
	if len(events) != 0 {
		t.Error("events is not 0")
		return
	}

	if _, err := repo.Insert(event1); err != nil {
		t.Error(err)
		return
	}

	if _, err := repo.Insert(event2); err != nil {
		t.Error(err)
		return
	}

	if _, err := repo.Insert(event3); err != nil {
		t.Error(err)
		return
	}

	events, err = repo.ListOverlapPairs()
	if err != nil {
		t.Error(err)
		return
	}

	if len(events) != 2 {
		t.Errorf("events is not 2 is %d", len(events))
		return
	}
}

func TestEventRepositoryGet(t *testing.T) {
	repo, err := getTestRepository()
	if err != nil {
		t.Error(err)
		return
	}
	event := getTestEvent()
	id, err := repo.Insert(event)
	if err != nil {
		t.Error(err)
		return
	}
	if id == 0 {
		t.Error("id is 0")
		return
	}

	event, err = repo.Get(id)
	if err != nil {
		t.Error(err)
		return
	}
	if event.ID != id {
		t.Error("id is not equal")
		return
	}
}

func TestEventRepositoryDelete(t *testing.T) {
	repo, err := getTestRepository()
	if err != nil {
		t.Error(err)
		return
	}
	event := getTestEvent()
	id, err := repo.Insert(event)
	if err != nil {
		t.Error(err)
		return
	}
	if id == 0 {
		t.Error("id is 0")
		return
	}

	if err := repo.Delete(id); err != nil {
		t.Error(err)
		return
	}

	event, err = repo.Get(id)
	if err == nil {
		t.Error("event is not nil")
		return
	}

	if event.ID != 0 {
		t.Error("event id is not 0")
		return
	}
}

func TestEventRepositoryUpdate(t *testing.T) {
	repo, err := getTestRepository()
	if err != nil {
		t.Error(err)
		return
	}
	event := getTestEvent()
	id, err := repo.Insert(event)
	if err != nil {
		t.Error(err)
		return
	}
	if id == 0 {
		t.Error("id is 0")
		return
	}

	event, err = repo.Get(id)
	if err != nil {
		t.Error(err)
		return
	}
	event.Title = "Test2"

	if err := repo.Update(event); err != nil {
		t.Error(err)
		return
	}

	event, err = repo.Get(id)
	if err != nil {
		t.Error(err)
		return
	}

	if event.Title != "Test2" {
		t.Error("event title is not equal")
		return
	}
}
