package lesson_group

import (
	"go-academy/business/lesson_group"
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

func (r *repository) InsertLessonGroup(lessonGroup lesson_group.LessonGroup) error {
	return r.db.Create(&lessonGroup).Error
}

func (r *repository) FindLessonTypeByID(id int) (lesson_type.LessonType, error) {
	var lessonType lesson_type.LessonType

	err := r.db.Find(&lessonType, id).Error

	return lessonType, err
}

func (r *repository) FindAllLessonGroup() (listLessonGroup []lesson_group.LessonGroup, err error) {
	var listLessonGroupDB []lesson_group.LessonGroup

	err = r.db.Find(&listLessonGroupDB).Error

	return listLessonGroupDB, err
}