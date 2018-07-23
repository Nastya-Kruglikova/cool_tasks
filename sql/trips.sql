CREATE TABLE trips(
  trip_id uuid DEFAULT uuid_generate_v1(),
  user_id uuid REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
  PRIMARY KEY (trip_id)
);
