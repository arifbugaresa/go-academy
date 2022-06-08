package lesson_group

import (
	"go-academy/api/v1/request"
	"go-academy/api/v1/response"
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

func (s *service) InsertLessonGroup(request request.InsertMultipleLessonGroup) (err error) {

	// convert lesson request to model
	listLessonGroup := s.convertDTOTOModelForMultipleInsert(request)

	// validasi apakah lesson tipe id ditemukan di database
	data, err := s.repository.FindLessonTypeByID(request.LessonTypeID)
	if err != nil {
		return business.ErrGetDataFromDB
	}

	if data.ID < 1 {
		return business.ErrDataNotFound
	}

	// insert multiple lesson group
	for _, data := range listLessonGroup {
		err := s.repository.InsertLessonGroup(data)
		if err != nil {
			return business.ErrInsertData
		}
	}

	return nil
}

func (s *service) convertDTOTOModelForMultipleInsert(dto request.InsertMultipleLessonGroup) (listLessonGroup []LessonGroup) {
	for _, data := range dto.LessonGroup {
		lessonGroup := LessonGroup{
			Name:         data.Name,
			Desc:         data.Desc,
			LessonTypeID: dto.LessonTypeID,
			Priority:     data.Priority,
			CreatedBy:    "",
		}

		listLessonGroup = append(listLessonGroup, lessonGroup)
	}

	return
}

/*
GET LIST LESSON GROUP
*/

func (s *service) FindAllLessonGroup() (response []response.GetListLessonGroup, err error) {
	listLessonGroupOnDB, err := s.repository.FindAllLessonGroup()
	if err != nil {
		return response, business.ErrGetDataFromDB
	}

	response = s.convertModelTODTOForGetList(listLessonGroupOnDB)

	return
}

func (s *service) convertModelTODTOForGetList(listLessonGroupOnDB []LessonGroup) (res []response.GetListLessonGroup) {
	for _, data := range listLessonGroupOnDB {
		lessonGroup := response.GetListLessonGroup{
			ID:           data.ID,
			Name:         data.Name,
			Desc:         data.Desc,
			LessonTypeID: data.LessonTypeID,
			CreatedAt:    data.CreatedAt,
			UpdatedAt:    data.UpdatedAt,
		}

		res = append(res, lessonGroup)
	}

	return
}
