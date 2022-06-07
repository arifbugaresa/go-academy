package lesson_type

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
			ID:           lessonType.ID,
			Name:         lessonType.Name,
			Desc:         lessonType.Desc,
			LastAccessed: lessonType.LastAccessed,
		}

		listLessonTypeResponse = append(listLessonTypeResponse, data)
	}

	return
}

/*
UPDATE LESSON TYPE
*/

func (s *service) UpdateLessonType(request request.UpdateLessonType) (err error) {

	// find data by id
	lessonTypeOnDB, err := s.FindLessonTypeByID(request.ID)
	if err != nil {
		return business.ErrInvalidBody
	}

	if lessonTypeOnDB.ID < 1 {
		return business.ErrDataNotFound
	}

	lessonTypeModel := s.convertDTOToModelForUpdate(request)

	return s.repository.UpdateLessonType(lessonTypeModel)
}

func (s *service) convertDTOToModelForUpdate(request request.UpdateLessonType) (lessonType LessonType) {
	return LessonType{
		ID:        request.ID,
		Name:      request.Name,
		Desc:      request.Desc,
		CreatedBy: "",
		UpdatedBy: "",
	}
}

/*
FIND LESSON TYPE BY ID
*/

func (s *service) FindLessonTypeByID(id int) (LessonType, error) {
	return s.repository.FindLessonTypeByID(id)
}
