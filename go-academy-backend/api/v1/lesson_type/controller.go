package lesson_type

import (
	"github.com/labstack/echo/v4"
	"go-academy/api/v1/common"
	"go-academy/api/v1/request"
	"go-academy/business/lesson_type"
)

type Controller struct {
	service lesson_type.Service
}

func NewController(service lesson_type.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) InsertLessonType(context echo.Context) error {
	var request request.InsertLessonType
	if err := context.Bind(&request); err != nil {
		return context.JSON(common.NewErrBindData())
	}

	dataFailed, isFailed := request.MandatoryValidation()
	if isFailed {
		return context.JSON(common.NewValidationResponse(dataFailed))
	}

	err := controller.service.InsertLessonType(request)
	if err != nil {
		return context.JSON(common.NewErrorBusinessResponse(err))
	}

	return context.JSON(common.NewSuccessResponseWithoutData())
}