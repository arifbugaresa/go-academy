package response

import "time"

type GetListLesson struct {
	ID            int       `json:"id"`
	LessonGroupID int       `json:"lesson_group_id"`
	Name          string    `json:"name"`
	Desc          string    `json:"desc"`
	Url           string    `json:"url"`
	Status        bool      `json:"status"`
	Priority      float64   `json:"priority"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
