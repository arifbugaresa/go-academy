package lesson

import (
	"go-academy/api/v1/common"
	"go-academy/api/v1/request"
	"go-academy/api/v1/response"
	"go-academy/business"
	"strconv"
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

/*
GET LIST ALL LESSON
*/

func (s *service) FindAllLesson() (response []response.GetListLesson, err error) {
	listLessonOnDB, err := s.repository.FindAllLesson()
	if err != nil {
		return response, business.ErrGetDataFromDB
	}

	response = s.convertModelTODTOForGetList(listLessonOnDB)

	return
}

func (s *service) convertModelTODTOForGetList(listLessonOnDB []Lesson) (res []response.GetListLesson) {
	for _, data := range listLessonOnDB {
		lesson := response.GetListLesson{
			ID:            data.ID,
			Name:          data.Name,
			Desc:          data.Desc,
			Url:           data.Url,
			Status:        data.Status,
			Priority:      data.Priority,
			LessonGroupID: data.LessonGroupID,
			CreatedAt:     data.CreatedAt,
			UpdatedAt:     data.UpdatedAt,
		}

		res = append(res, lesson)
	}

	return
}

/*
FIND BY ID FOR SEARCH
*/

func (s *service) FindLessonByID(id string) (lessonOnDB Lesson, err error) {
	lessonOnDB, err = s.repository.FindLessonByID(id)
	if err != nil {
		return lessonOnDB, business.ErrGetDataFromDB
	}

	if lessonOnDB.ID < 1 {
		return lessonOnDB, business.ErrDataNotFound
	}

	return
}

/*
DELETE LESSON
*/

func (s *service) DeleteLessonByID(id string) (err error) {
	lessonOnDB, err := s.FindLessonByID(id)
	if err != nil {
		return
	}

	return s.repository.DeleteLessonByID(strconv.Itoa(lessonOnDB.ID))
}
