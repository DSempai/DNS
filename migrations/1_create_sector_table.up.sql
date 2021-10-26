CREATE TABLE IF NOT EXISTS sectors (
    id serial,
    sector_id INTEGER PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    active BOOLEAN NOT NULL
);

-- Fill database for testing purposes.
INSERT INTO sectors(sector_id, active) VALUES(1, true);