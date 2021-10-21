CREATE TABLE IF NOT EXISTS sectors (
    id serial,
    sector_id INTEGER PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    active BOOLEAN NOT NULL
);