package entities

import "time"

type ExampleEntity struct {
	Id         uint      `gorm:"primaryKey;autoIncrement"`
	Name       string    `gorm:"type:varchar(100);not null"`
	Created_at time.Time `gorm:"autoCreateTime"`
	Updated_at time.Time `gorm:"autoUpdateTime"`
	Count      int       `gorm:"type:int;not null"`
}

func (ExampleEntity) TableName() string {
	return "examples"
}
