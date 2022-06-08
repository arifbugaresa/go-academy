package lesson

import (
	"go-academy/api/v1/common"
	"go-academy/api/v1/request"
	"go-academy/business"
)

type service struct {
	repository Repository
}

// NewService is used to inject repo product to service
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

/*
INSERT LESSON GROUP
*/

func (s *service) InsertLesson(request request.InsertMultipleLesson) (err error) {
	funcName := "InsertLesson"

	// convert lesson request to model
	listLessonGroup := s.convertDTOTOModelForMultipleInsert(request)

	// validasi apakah lesson tipe id ditemukan di database
	data, err := s.repository.FindLessonGroupByID(request.LessonGroupID)
	if err != nil {
		return business.ErrGetDataFromDB
	}

	if data.ID < 1 {
		return business.GenerateErrDataNotFound(funcName, common.LESSON, common.LESSON, request.LessonGroupID)
	}

	// insert multiple lesson group
	for _, data := range listLessonGroup {
		err := s.repository.InsertLesson(data)
		if err != nil {
			return business.ErrInsertData
		}
	}

	return nil
}

func (s *service) convertDTOTOModelForMultipleInsert(dto request.InsertMultipleLesson) (listLesson []Lesson) {
	for _, data := range dto.Lesson {
		lessonGroup := Lesson{
			Name:          data.Name,
			Desc:          data.Desc,
			LessonGroupID: dto.LessonGroupID,
			Priority:      data.Priority,
			Url:           data.Url,
			CreatedBy:     "",
		}

		listLesson = append(listLesson, lessonGroup)
	}

	return
}