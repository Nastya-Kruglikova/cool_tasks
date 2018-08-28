package events_test

import (
	"bytes"
	"fmt"
	"github.com/Nastya-Kruglikova/cool_tasks/src/model"
	"github.com/Nastya-Kruglikova/cool_tasks/src/service"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var router = service.NewRouter()

type EventsTestCase struct {
	name            string
	url             string
	want            int
	mockedGetEvents []model.Event
	testDataId      string
	testDataEv      string
	mockedEventsErr error
}

func TestGetByRequestHandler(t *testing.T) {
	tests := []EventsTestCase{
		{
			name:            "Get_Events_200",
			url:             "/v1/events",
			want:            200,
			mockedGetEvents: []model.Event{},
			mockedEventsErr: nil,
		},
		{
			name:            "Get_Events_404",
			url:             "/v1/events?mock=890",
			want:            404,
			mockedGetEvents: []model.Event{},
			mockedEventsErr: http.ErrBodyNotAllowed,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model.MockedGetEvents(tc.mockedGetEvents, tc.mockedEventsErr)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)
			router.ServeHTTP(rec, req)
			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}

func TestAddToTripHandler(t *testing.T) {
	tests := []EventsTestCase{
		{
			name:            "Add_Events_201",
			url:             "/v1/events",
			want:            201,
			testDataId:      "00000000-0000-0000-0000-000000000001",
			testDataEv:      "00000000-0000-0000-0000-000000000001",
			mockedEventsErr: nil,
		},
		{
			name:            "Add_Events_400",
			url:             "/v1/events",
			want:            400,
			testDataId:      "00000000-0000-0000-0000-000000000001",
			testDataEv:      "asdsad",
			mockedEventsErr: nil,
		},
		{
			name:            "Add_Events_400_2",
			url:             "/v1/events",
			want:            400,
			testDataId:      "asdasd",
			testDataEv:      "00000000-0000-0000-0000-000000000001",
			mockedEventsErr: nil,
		},
		{
			name:            "Add_Events_400_3",
			url:             "/v1/events",
			want:            400,
			testDataId:      "00000000-0000-0000-0000-000000000001",
			testDataEv:      "00000000-0000-0000-0000-000000000001",
			mockedEventsErr: error(new(http.ProtocolError)),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data := url.Values{}
			data.Add("event_id", tc.testDataEv)
			data.Add("trip_id", tc.testDataId)
			model.MockedAddEventToTrip(tc.mockedEventsErr)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewBufferString(data.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
			router.ServeHTTP(rec, req)
			fmt.Println(rec.Code)
			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
			data.Del("event_id")
			data.Del("trip_id")
		})
	}
}

func TestGetByTripHandler(t *testing.T) {
	tests := []EventsTestCase{
		{
			name:            "Get_Events_200",
			url:             "/v1/events/trip/00000000-0000-0000-0000-000000000001",
			want:            200,
			mockedGetEvents: []model.Event{},
			mockedEventsErr: nil,
		},
		{
			name:            "Get_Events_400",
			url:             "/v1/events/trip/sadsad",
			want:            400,
			mockedGetEvents: []model.Event{},
			mockedEventsErr: nil,
		},
		{
			name:            "Get_Events_404",
			url:             "/v1/events/trip/00000000-0000-0000-0000-000000000009",
			want:            404,
			mockedGetEvents: []model.Event{},
			mockedEventsErr: http.ErrLineTooLong,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model.MockedGetEventsByTrip(tc.mockedGetEvents, tc.mockedEventsErr)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)
			fmt.Println(rec.Code)
			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}