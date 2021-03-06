package lesson_type

import "time"

type LessonType struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Desc         string
	LastAccessed time.Time `sql:"DEFAULT:'current_timestamp'"`
	CreatedAt    time.Time `sql:"DEFAULT:'current_timestamp'"`
	UpdatedAt    time.Time `sql:"DEFAULT:'current_timestamp'"`
	CreatedBy    string
	UpdatedBy    string
}