package lesson_type

import (
	"go-academy/api/v1/request"
	"go-academy/api/v1/response"
)

type Service interface {
	InsertLessonType(request request.InsertLessonType) error
	GetListLessonType() ([]response.GetListLessonType, error)
}

type Repository interface {
	InsertLessonType(lessonType LessonType) error
	GetListLessonType() ([]LessonType, error)
}
