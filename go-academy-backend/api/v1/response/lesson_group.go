package response

import "time"

type GetListLessonGroup struct {
	ID           int
	Name         string
	Desc         string
	LessonTypeID int
	Priority     float64
	CreatedAt    time.Time `sql:"DEFAULT:'current_timestamp'"`
	UpdatedAt    time.Time `sql:"DEFAULT:'current_timestamp'"`
}