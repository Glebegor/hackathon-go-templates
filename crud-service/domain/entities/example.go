package entities

import "time"

type ExampleEntity struct {
	Id         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string    `json:"name" gorm:"type:varchar(100);not null"`
	Created_at time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Count      int       `json:"count" gorm:"type:int;not null"`
}

func (ExampleEntity) TableName() string {
	return "examples"
}
