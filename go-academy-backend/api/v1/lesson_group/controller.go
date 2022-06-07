package lesson_group

import (
	"github.com/labstack/echo/v4"
	"go-academy/api/v1/common"
	"go-academy/api/v1/request"
	"go-academy/business/lesson_group"
)

type Controller struct {
	service lesson_group.Service
}

func NewController(service lesson_group.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) InsertLessonGroup(context echo.Context) error {
	var request request.InsertMultipleLessonGroup
	if err := context.Bind(&request); err != nil {
		return context.JSON(common.NewErrBindData())
	}

	dataFailed, isFailed := request.MandatoryValidation()
	if isFailed {
		return context.JSON(common.NewValidationResponse(dataFailed))
	}

	err := controller.service.InsertLessonGroup(request)
	if err != nil {
		return context.JSON(common.NewErrorBusinessResponse(err))
	}

	return context.JSON(common.NewSuccessResponseWithoutData(common.SUCCESS_INSERT_DATA))
}
