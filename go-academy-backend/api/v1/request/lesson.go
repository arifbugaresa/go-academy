package request

import "go-academy/api/v1/common"

type InsertMultipleLesson struct {
	LessonGroupID int            `json:"lesson_group_id"`
	Lesson        []InsertLesson `json:"lesson"`
}

type InsertLesson struct {
	Name     string  `json:"name"`
	Desc     string  `json:"desc"`
	Url      string  `json:"url"`
	Status   bool    `json:"status"`
	Priority float64 `json:"priority"`
}

func (dto InsertMultipleLesson) MandatoryValidation() (dataFailed string, error bool) {

	if common.IsNumericZeroOrMinus(dto.LessonGroupID) {
		return common.LESSON_TYPE_ID, true
	}

	for _, data := range dto.Lesson {
		if common.IsStringEmpty(data.Name) {
			return common.NAME, true
		}

		if common.IsFloatZeroOrMinus(data.Priority) {
			return common.PRIORITY, true
		}

		if common.IsStringEmpty(data.Url) {
			return common.URL, true
		}
	}

	return
}
