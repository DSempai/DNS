CREATE TABLE IF NOT EXISTS sectors (
    id serial,
    sector_id INTEGER PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    active BOOLEAN NOT NULL
);