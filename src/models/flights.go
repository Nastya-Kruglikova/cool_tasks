package models

import (
	"errors"
	sq "github.com/Masterminds/squirrel"
	. "github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"net/url"
	"time"
)

const (
	addFlightToTrip = "INSERT INTO trips_flights (flight_id, trip_id) VALUES ($1, $2)"
	getFlightByTrip = "SELECT flights.* FROM flights INNER JOIN trips_flights ON flights.id=trips_flights.flight_id AND trips_flights.trip_id=$1"
)

type Flight struct {
	ID            uuid.UUID
	DepartureCity string
	DepartureTime time.Time
	DepartureDate time.Time
	ArrivalCity   string
	ArrivalTime   time.Time
	ArrivalDate   time.Time
	Price         int
}

var AddFlightToTrip = func(flightID uuid.UUID, tripID uuid.UUID) error {
	_, err := DB.Exec(addFlightToTrip, flightID, tripID)
	return err
}

var GetFlightsByTrip = func(tripID uuid.UUID) ([]Flight, error) {
	rows, err := DB.Query(getFlightByTrip, tripID)
	if err != nil {
		return nil, err
	}

	flights := make([]Flight, 0)

	for rows.Next() {
		var f Flight
		if err := rows.Scan(&f.ID, &f.DepartureCity, &f.DepartureTime, &f.DepartureDate, &f.ArrivalCity, &f.ArrivalDate, &f.ArrivalTime, &f.Price); err != nil {
			return nil, err
		}
		flights = append(flights, f)
	}
	return flights, nil
}

var GetFlightsByRequest = func(params url.Values) ([]Flight, error) {

	var and sq.And = nil
	flights := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select("*").From("flights")
	for key, value := range params {
		var or sq.Or = nil
		switch key {
		case "departure_city", "arrival_city":
			if len(value) > 1 {
				for _, v := range value {
					or = append(or, sq.Eq{key: v})
				}
				and = append(and, or)
			} else {
				and = append(and, sq.Eq{key: value[0]})
			}
		case "departure_time", "departure_date", "arrival_time", "arrival_date", "price":
			if len(value) > 1 {
				and = append(and, sq.And{sq.GtOrEq{key: value[1]}, sq.LtOrEq{key: value[0]}})
			} else {
				and = append(and, sq.Eq{key: value[0]})
			}
		case "id":
			and = append(and, sq.Eq{key: value[0]})
		default:
			return nil, errors.New("ERROR: Bad request")
		}
	}

	req := flights.Where(and)

	request, args, err := req.ToSql()
	if err != nil {
		return nil, errors.New("ERROR: Bad request")
	}
	rows, err := DB.Query(request, args...)
	if err != nil {
		return nil, err
	}

	flightsArray := make([]Flight, 0)

	for rows.Next() {
		var f Flight
		if err := rows.Scan(&f.ID, &f.DepartureCity, &f.DepartureTime, &f.DepartureDate, &f.ArrivalCity, &f.ArrivalDate, &f.ArrivalTime, &f.Price); err != nil {
			return nil, err
		}
		flightsArray = append(flightsArray, f)
	}
	return flightsArray, nil
}
