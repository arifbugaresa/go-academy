package lesson_type

import (
	"go-academy/api/v1/request"
	"go-academy/api/v1/response"
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
INSERT LESSON TYPE
*/

func (s *service) InsertLessonType(request request.InsertLessonType) (err error) {

	// convert lesson type request to model
	lessonType := s.convertDTOtoModelForInsert(request)

	// insert into database
	err = s.repository.InsertLessonType(lessonType)
	if err != nil {
		return
	}

	return
}

func (s *service) convertDTOtoModelForInsert(dto request.InsertLessonType) LessonType {
	return LessonType{
		Name:      dto.Name,
		Desc:      dto.Desc,
		CreatedBy: "",
		UpdatedBy: "",
	}
}

/*
GET LIST LESSON TYPE
*/

func (s *service) GetListLessonType() (listLessonType []response.GetListLessonType, err error) {
	listLessonTypeDB, err := s.repository.GetListLessonType()
	if err != nil {
		return
	}

	listLessonType = s.convertModelToDTOForGetList(listLessonTypeDB)

	return
}

func (s *service) convertModelToDTOForGetList(listLessonTypeDB []LessonType) (listLessonTypeResponse []response.GetListLessonType) {
	for _, lessonType := range listLessonTypeDB {
		data := response.GetListLessonType{
			ID:   lessonType.ID,
			Name: lessonType.Name,
			Desc: lessonType.Desc,
		}

		listLessonTypeResponse = append(listLessonTypeResponse, data)
	}

	return
}
