package lesson_group

import (
	"go-academy/api/v1/request"
	"go-academy/api/v1/response"
	"go-academy/business/lesson_type"
)

type Service interface {
	InsertLessonGroup(request request.InsertMultipleLessonGroup) (err error)
	FindAllLessonGroup() (response []response.GetListLessonGroup, err error)
	FindLessonGroupByID(id string) (response LessonGroup, err error)
	DeleteLessonGroupByID(id string) (err error)
}

type Repository interface {
	InsertLessonGroup(lessonGroup LessonGroup) error
	FindLessonTypeByID(id int) (lessonType lesson_type.LessonType, err error)
	FindAllLessonGroup() (listLessonGroup []LessonGroup, err error)
	FindLessonGroupByID(id string) (lessonGroupo LessonGroup, err error)
	DeleteLessonGroupByID(id string) (err error)
	DeleteLessonByLessonID(id string) (err error)
}
