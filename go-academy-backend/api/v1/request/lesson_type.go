package request

import "go-academy/api/v1/common"

type InsertLessonType struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (dto InsertLessonType) MandatoryValidation() (dataFailed string, error bool) {

	if common.IsStringEmpty(dto.Name) && common.IsStringEmpty(dto.Desc) {
		return "request", true
	}

	if common.IsStringEmpty(dto.Name){
		return "name", true
	}

	if common.IsStringEmpty(dto.Desc) {
		return "desc", true
	}

	return
}