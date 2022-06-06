package lesson_type

import "go-academy/api/v1/request"

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
func (s *service) InsertLessonType(request request.InsertLessonType) error {

	// convert lesson type request to model
	lessonType := s.convertDTOtoModelForInsert(request)

	// insert into database
	err := s.repository.InsertLessonType(lessonType)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) convertDTOtoModelForInsert(dto request.InsertLessonType) LessonType {
	return LessonType{
		Name:      dto.Name,
		Desc:      dto.Desc,
		CreatedBy: "",
		UpdatedBy: "",
	}
}
