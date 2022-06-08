package lesson

import (
	"github.com/labstack/echo/v4"
	"go-academy/api/v1/common"
	"go-academy/api/v1/request"
	"go-academy/business/lesson"
)

type Controller struct {
	service lesson.Service
}

func NewController(service lesson.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) InsertLesson(context echo.Context) error {
	var request request.InsertMultipleLesson
	if err := context.Bind(&request); err != nil {
		return context.JSON(common.NewErrBindData())
	}

	dataFailed, isFailed := request.MandatoryValidation()
	if isFailed {
		return context.JSON(common.NewValidationResponse(dataFailed))
	}

	err := controller.service.InsertLesson(request)
	if err != nil {
		return context.JSON(common.NewErrorBusinessResponse(err))
	}

	return context.JSON(common.NewSuccessResponseWithoutData(common.SUCCESS_INSERT_DATA))
}

func (controller *Controller) FindAllLesson(context echo.Context) (err error) {
	listLesson, err := controller.service.FindAllLesson()
	if err != nil {
		return context.JSON(common.NewErrorBusinessResponse(err))
	}

	return context.JSON(common.NewSuccessResponse(common.SUCCESS_GET_LIST_DATA, listLesson))
}