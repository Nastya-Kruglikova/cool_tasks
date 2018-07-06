package museums_test

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services"
	"github.com/Nastya-Kruglikova/cool_tasks/src/museums"
	"github.com/satori/go.uuid"
	"net/url"
	"bytes"
)

var router = services.NewRouter()

type MuseumsTestCase struct {
	name             string
	url              string
	want             int
	mockedGetMuseums []museums.Museum
}

func TestGetMuseumsHandler(t *testing.T) {
	tests := []MuseumsTestCase{
		{
			name:             "Get_Museums_200",
			url:              "/v1/museums",
			want:             200,
			mockedGetMuseums: []museums.Museum{},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			museums.MockedGetMuseums()
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}

func TestGetMuseumByCityHandler(t *testing.T) {
	MuseumId, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")
	tests := []MuseumsTestCase{
		{
			name: "Get_Museums_200",
			url:  "/v1/museums/Paris",
			want: 200,
			mockedGetMuseums: []museums.Museum{
				{
					MuseumId,
					"Louvre",
					"Paris",
					1111,
					1,
					2,
					"History",
					"Cool",
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			museums.MockedGetMuseumByCity()
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}

func TestAddMuseumToTripHandler(t *testing.T) {
	tests := []MuseumsTestCase{
		{
			name: "Add_Museum_200",
			url:  "/v1/museums",
			want: 200,
		},
	}
	data := url.Values{}
	data.Add("museum", "00000000-0000-0000-0000-000000000001")
	data.Add("trip", "00000000-0000-0000-0000-000000000001")
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			museums.MockedAddMuseum()
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewBufferString(data.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}

func TestGetMuseumByTripHandler(t *testing.T) {
	tests := []MuseumsTestCase{
		{
			name:             "Get_Museums_200",
			url:              "/v1/museums/trip/00000000-0000-0000-0000-000000000001",
			want:             200,
			mockedGetMuseums: []museums.Museum{},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			museums.MockedGetMuseumsByTrip()
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}