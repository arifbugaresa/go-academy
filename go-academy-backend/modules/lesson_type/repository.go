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

func (r *repository) GetListLessonType() ([]lesson_type.LessonType, error) {
	var listLessonType []lesson_type.LessonType
	err := r.db.Find(&listLessonType).Error

	return listLessonType, err
}

func (r *repository) UpdateLessonType(lt lesson_type.LessonType) error {
	return r.db.Model(&lt).Updates(
		lesson_type.LessonType{
			Name:         lt.Name,
			UpdatedAt:    lt.UpdatedAt,
			LastAccessed: lt.LastAccessed,
			Desc:         lt.Desc,
		},
	).Error
}

func (r *repository) FindLessonTypeByID(id int) (lesson_type.LessonType, error) {
	var lessonType lesson_type.LessonType

	err := r.db.Find(&lessonType, id).Error

	return lessonType, err
}
