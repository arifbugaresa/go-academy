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
