package response

import "time"

type GetListLessonGroup struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Desc         string    `json:"desc"`
	LessonTypeID int       `json:"lesson_type_id"`
	Priority     float64   `json:"priority"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
