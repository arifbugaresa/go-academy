package lesson_group

import "time"

type LessonGroup struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Desc         string
	LessonTypeID int
	Priority     float64
	CreatedAt    time.Time `sql:"DEFAULT:'current_timestamp'"`
	UpdatedAt    time.Time `sql:"DEFAULT:'current_timestamp'"`
	CreatedBy    string
	UpdatedBy    string
}
