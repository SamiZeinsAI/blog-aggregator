-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    time_created TIMESTAMP,
    time_updated TIMESTAMP,
    name TEXT,
    url TEXT UNIQUE,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE
);
-- +goose Down
DROP TABLE feeds;