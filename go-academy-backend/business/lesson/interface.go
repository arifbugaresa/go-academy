package lesson

import (
	"go-academy/api/v1/request"
	"go-academy/business/lesson_group"
)

type Service interface {
	InsertLesson(request request.InsertMultipleLesson) (err error)
	//FindAllLessonGroup() (response []response.GetListLessonGroup, err error)
	//FindLessonGroupByID(id string) (response LessonGroup, err error)
	//DeleteLessonGroupByID(id string) (err error)
}

type Repository interface {
	InsertLesson(lesson Lesson) error
	FindLessonGroupByID(id int) (lessonGroup lesson_group.LessonGroup, err error)
	//FindAllLessonGroup() (listLessonGroup []LessonGroup, err error)
	//FindLessonGroupByID(id string) (lessonGroupo LessonGroup, err error)
	//DeleteLessonByID(id string) (err error)
}

