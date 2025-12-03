package entities

import "time"

type ExampleEntity struct {
	Id         int       `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
	Count      int       `json:"count" db:"count"`
}
