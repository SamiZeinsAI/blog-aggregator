// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	ID          uuid.UUID
	TimeCreated sql.NullTime
	TimeUpdated sql.NullTime
	Name        sql.NullString
	Url         sql.NullString
	UserID      uuid.NullUUID
}

type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}
