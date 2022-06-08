package lesson

import (
	"go-academy/api/v1/request"
	"go-academy/api/v1/response"
	"go-academy/business/lesson_group"
)

type Service interface {
	InsertLesson(request request.InsertMultipleLesson) (err error)
	FindAllLesson() (response []response.GetListLesson, err error)
	FindLessonByID(id string) (lessonOnDB Lesson, err error)
	DeleteLessonByID(id string) (err error)
}

type Repository interface {
	InsertLesson(lesson Lesson) error
	FindLessonGroupByID(id int) (lessonGroup lesson_group.LessonGroup, err error)
	FindAllLesson() (listLesson []Lesson, err error)
	FindLessonByID(id string) (res Lesson, err error)
	DeleteLessonByID(id string) (err error)
}

