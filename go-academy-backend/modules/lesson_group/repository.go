package lesson_group

import (
	"go-academy/business/lesson"
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

func (r *repository) DeleteLessonGroupByID(id string) (err error) {
	return r.db.Delete(&lesson_group.LessonGroup{}, id).Error
}

func (r *repository) DeleteLessonByLessonID(id string) (err error) {
	return r.db.Where("lesson_group_id = ?", id).Delete(&lesson.Lesson{}).Error
}

func (r *repository) FindLessonGroupByID(id string) (res lesson_group.LessonGroup, err error) {
	var lessonGroup lesson_group.LessonGroup
	err = r.db.Find(&lessonGroup, id).Error
	return lessonGroup, err
}
