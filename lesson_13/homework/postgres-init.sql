CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    second_name TEXT NOT NULL,
    birthdate DATE NOT NULL,
    city TEXT NOT NULL
);

COPY users (second_name, first_name, birthdate, city)
FROM '/var/lib/postgresql/data/pgdata/people.csv' CSV;
