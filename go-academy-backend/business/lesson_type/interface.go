package lesson_type

import (
	"go-academy/api/v1/request"
	"go-academy/api/v1/response"
)

type Service interface {
	InsertLessonType(request.InsertLessonType) error
	GetListLessonType() ([]response.GetListLessonType, error)
	UpdateLessonType(request.UpdateLessonType) error
	DeleteLessonType(id int) (err error)
}

type Repository interface {
	InsertLessonType(LessonType) error
	GetListLessonType() ([]LessonType, error)
	UpdateLessonType(LessonType) error
	FindLessonTypeByID(int) (LessonType, error)
	DeleteLessonTypeByID(int) error
}
