create extension if not exists citext;


-- CREATE ALL TABLES

CREATE TABLE IF NOT EXISTS airports
(
    id       BIGSERIAL    NOT NULL CONSTRAINT airports_pk PRIMARY KEY,
    address TEXT  NOT NULL ,
    "name" TEXT  NOT NULL ,
);

CREATE TABLE IF NOT EXISTS aircrafts
(
    id       BIGSERIAL    NOT NULL CONSTRAINT airports_pk PRIMARY KEY,
    model TEXT  NOT NULL ,
    "name" TEXT  NOT NULL ,
);

CREATE TABLE IF NOT EXISTS seats
(
    id       BIGSERIAL    NOT NULL CONSTRAINT seats_pk PRIMARY KEY,
    seat_class TEXT  NOT NULL ,
    aircraft_id INT  NOT NULL ,
    FOREIGN KEY (aircraft_id) REFERENCES aircrafts (id),
);



CREATE TABLE IF NOT EXISTS flights
(
    id       BIGSERIAL    NOT NULL CONSTRAINT flights_pk PRIMARY KEY,
    departure_id INT  NOT NULL ,
    arrival_id INT  NOT NULL ,
    avg_duration_minutes INT NOT NULL,

    FOREIGN KEY (departure_id) REFERENCES airports (id),
    FOREIGN KEY (arrival_id) REFERENCES airports (id),
);

CREATE TABLE IF NOT EXISTS schedules
(
    id       BIGSERIAL    NOT NULL CONSTRAINT schedules_pk PRIMARY KEY,
    date TIMESTAMPTZ  NOT NULL,
    flight_id INT NOT NULL ,

    FOREIGN KEY (flight_id) REFERENCES flights (id),
);

CREATE TABLE IF NOT EXISTS flights_instances
(
    id       BIGSERIAL    NOT NULL CONSTRAINT flights_instances_pk PRIMARY KEY,
    flight_id INT NOT NULL ,
    aircraft_id INT  NOT NULL ,
    status TEXT  NOT NULL ,

    FOREIGN KEY (flight_id) REFERENCES flights (id),
    FOREIGN KEY (aircraft_id) REFERENCES aircrafts (id),

);

CREATE TABLE IF NOT EXISTS flights_instances_seats
(
    id       BIGSERIAL    NOT NULL CONSTRAINT flights_instances_pk PRIMARY KEY,
    flight_instance_id INT NOT NULL ,
    seat_id INT  NOT NULL ,
    cost INT  NOT NULL ,

    FOREIGN KEY (flight_instance_id) REFERENCES flights_instances (id),
    FOREIGN KEY (seat_id) REFERENCES seats (id),

);


CREATE TABLE IF NOT EXISTS flight_reservation
(
    id       BIGSERIAL    NOT NULL CONSTRAINT flight_reservation_pk PRIMARY KEY,
    flight_instance_id INT NOT NULL ,

    FOREIGN KEY (flight_instance_id) REFERENCES flights_instances (id),
);

CREATE TABLE IF NOT EXISTS flight_reservation_ticket_info
(
    id       BIGSERIAL    NOT NULL CONSTRAINT flight_reservation_ticket_info_pk PRIMARY KEY,
    flight_reservation_id INT NOT NULL ,
    include_baggage BOOLEAN  NOT NULL ,
    passenger_id            NOT NULL,

    FOREIGN KEY (flight_reservation_id) REFERENCES flight_reservation (id),
);

