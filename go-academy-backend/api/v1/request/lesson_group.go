package request

import "go-academy/api/v1/common"

type InsertMultipleLessonGroup struct {
	LessonTypeID int                 `json:"lesson_type_id"`
	LessonGroup  []InsertLessonGroup `json:"lesson_group"`
}

type InsertLessonGroup struct {
	Name     string  `json:"name"`
	Desc     string  `json:"desc"`
	Priority float64 `json:"priority"`
}

func (dto InsertMultipleLessonGroup) MandatoryValidation() (dataFailed string, error bool) {

	if common.IsNumericZeroOrMinus(dto.LessonTypeID) {
		return "lesson_type_id", true
	}

	for _, data := range dto.LessonGroup {
		if common.IsStringEmpty(data.Name) {
			return "name", true
		}

		if common.IsFloatZeroOrMinus(data.Priority) {
			return "name", true
		}
	}

	return
}
