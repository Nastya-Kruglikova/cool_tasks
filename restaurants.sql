-- RESTAURANTS

CREATE TABLE restaurants (
   id uuid DEFAULT uuid_generate_v1(),
   name VARCHAR(100) NOT NULL,
   location VARCHAR(100) NOT NULL,
   stars int,
   prices int,
   description TEXT,
   PRIMARY KEY (id)
);

INSERT INTO restaurants
  (id,  name, location, stars, prices, description)
  VALUES ( uuid_generate_v1(), 'Крива липа', 'Lviv', 4, 3, 'Culinary studio "Krivaya Lipa" is an author's cuisine without GMOs. True culinary masterpieces of only the best quality and fresh produce from well-known masters.');
INSERT INTO restaurants
  (id,  name, location, stars, prices, description)
  VALUES ( uuid_generate_v1(), 'Криівка', 'Lviv', 5, 5, 'The authentic institution, decorated in the form of a field shelter UPA, is in the basement of one of the houses');
INSERT INTO restaurants
  (id,  name, location, stars, prices, description)
  VALUES ( uuid_generate_v1(), 'Живий хліб', 'Lviv', 5, 3, 'Bread and rolls are cooked here on natural starter from Italian flour. For croissants, use French butter.');
INSERT INTO restaurants
  (id,  name, location, stars, prices, description)
  VALUES ( uuid_generate_v1(), 'Фан-бар Банка', 'Lviv', 5, 4, 'Conceptual Democratic Bar, where for the first time in Ukraine all meals and drinks are served exclusively in traditional glass jars. Everything should be in the cans - bottles, plates, glasses and other similar dishes are forbidden in the bar.');


-- TRIP_RESTAURANTS

CREATE TABLE trips_restaurants(
  id uuid DEFAULT uuid_generate_v1(),
  trip_id uuid REFERENCES trips (trip_id) ON DELETE CASCADE,
  restaurant_id uuid REFERENCES restaurants (id) ON DELETE CASCADE,
  PRIMARY KEY (id)
);
