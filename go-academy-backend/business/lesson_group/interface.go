package lesson_group

import (
	"go-academy/api/v1/request"
	"go-academy/business/lesson_type"
)

type Service interface {
	InsertLessonGroup(request request.InsertMultipleLessonGroup) (err error)
}

type Repository interface {
	InsertLessonGroup(lessonGroup LessonGroup) error
	FindLessonTypeByID(id int) (lessonType lesson_type.LessonType, err error)
}
