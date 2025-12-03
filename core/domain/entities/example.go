package entities

import "time"

type ExampleEntity struct {
	id         int       `json:"id" db:"id"`
	name       string    `json:"name" db:"name"`
	created_at time.Time `json:"created_at" db:"created_at"`
	updated_at time.Time `json:"updated_at" db:"updated_at"`
	count      int       `json:"count" db:"count"`
}
