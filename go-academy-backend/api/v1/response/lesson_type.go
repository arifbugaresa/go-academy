package response

import "time"

type GetListLessonType struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Desc         string    `json:"desc"`
	LastAccessed time.Time `json:"last_accessed"`
}
