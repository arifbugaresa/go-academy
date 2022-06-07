package lesson

import "time"

type Lesson struct {
	ID            int `gorm:"primaryKey"`
	Name          string
	Desc          string
	LessonGroupID int
	Status        bool
	Priority      float64
	CreatedAt     time.Time `sql:"DEFAULT:'current_timestamp'"`
	UpdatedAt     time.Time `sql:"DEFAULT:'current_timestamp'"`
	CreatedBy     string
	UpdatedBy     string
}
