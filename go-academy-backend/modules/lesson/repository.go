package lesson

import (
	"go-academy/business/lesson"
	"go-academy/business/lesson_group"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db,
	}
}

func (r *repository) InsertLesson(lesson lesson.Lesson) (err error) {
	return r.db.Create(&lesson).Error
}

func (r *repository) FindLessonGroupByID(id int) (lessonGroup lesson_group.LessonGroup, err error) {

	err = r.db.Find(&lessonGroup, id).Error

	return
}

func (r *repository) FindAllLesson() (listLesson []lesson.Lesson, err error) {

	err = r.db.Find(&listLesson).Error

	return
}

func (r *repository) FindLessonByID(id string) (res lesson.Lesson, err error) {
	var lesson lesson.Lesson
	err = r.db.Find(&lesson, id).Error
	return lesson, err
}

func (r *repository) DeleteLessonByID(id string) (err error) {
	return r.db.Delete(lesson.Lesson{}, id).Error
}

