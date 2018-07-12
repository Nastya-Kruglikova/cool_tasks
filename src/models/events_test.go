package models

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"net/http"
	"testing"
	"time"
)

func TestGetEventsByRequest(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, err = sqlmock.New()
	defer func() { database.DB = originalDB }()
	testTime, _ := time.Parse("15:04:05", "12:00:00")
	EventId, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	var expects = []Event{
		{
			ID:       EventId,
			Title:    "Careerday",
			Category: "work",
			Town:     "Kiev",
			Date:     testTime,
			Price:    50,
		},
		{
			ID:       EventId,
			Title:    "ProjectX",
			Category: "entertaiment",
			Town:     "Lviv",
			Date:     testTime,
			Price:    300,
		},
	}

	rows := sqlmock.NewRows([]string{"ID", "Title", "Category", "Town", "Date", "Price"}).
		AddRow(EventId.Bytes(), "Careerday", "work", "Kiev", testTime, 50).
		AddRow(EventId.Bytes(), "ProjectX", "entertaiment", "Lviv", testTime, 300)

	mock.ExpectQuery("SELECT (.+) FROM events").WillReturnRows(rows)

	req, _ := http.NewRequest(http.MethodGet, "/v1/events", nil)

	result, err := GetEventsByRequest(req.URL.Query())

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	for i := 0; i < len(result); i++ {
		if expects[i] != result[i] {
			t.Error("Expected:", expects[i], "Was:", result[i])
		}
	}
}

func TestAddEventToTrip(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, err = sqlmock.New()
	defer func() { database.DB = originalDB }()

	TripId, _ := uuid.FromString("00000000-0000-0000-0000-000000000002")
	EventId, _ := uuid.FromString("00000000-0000-0000-0000-000000000003")

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec("INSERT INTO trips_events").WithArgs(EventId, TripId).WillReturnResult(sqlmock.NewResult(1, 1))
	if err := AddEventToTrip(EventId, TripId); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetEventsByTrip(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, err = sqlmock.New()
	defer func() { database.DB = originalDB }()

	testTime, _ := time.Parse("15:04:05", "12:00:00")
	EventId, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expected := Event{

		EventId,
		"Careerday",
		"work",
		"Kiev",
		testTime,
		50,
	}

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"EventId", "Title", "Category", "Town", "Date", "Price"}).
		AddRow(EventId.Bytes(), "Careerday", "work", "Kiev", testTime, 50)

	mock.ExpectQuery("SELECT (.+) FROM events INNER JOIN trips_events ON events.id=trips_events.event_id AND trips_events.trip_id=\\$1").WithArgs(expected.ID).WillReturnRows(rows)

	result, err := GetEventsByTrip(EventId)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if expected != result[0] {
		t.Error("Expected:", expected, "Was:", result[0])
	}
}