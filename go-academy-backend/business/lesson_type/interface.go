package lesson_type

import "go-academy/api/v1/request"

type Service interface {
	InsertLessonType(request request.InsertLessonType ) error
}

type Repository interface {
	InsertLessonType(lessonType LessonType) error
}
