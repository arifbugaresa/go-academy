package lesson_type

import (
	"go-academy/business/lesson_type"
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

func (r *repository) InsertLessonType(lessonType lesson_type.LessonType) error {
	return r.db.Create(&lessonType).Error
}