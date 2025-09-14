package model

import "time"

type Payment struct {
	ID        uint64    `db:"id"`
	Name      string    `db:"name"`
	Type      string    `db:"type"`
	Logo      string    `db:"logo"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
