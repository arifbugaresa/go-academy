package lesson_type

import (
	"github.com/labstack/echo/v4"
	"go-academy/api/v1/common"
	"go-academy/api/v1/request"
	"go-academy/api/v1/response"
	"go-academy/business/lesson_type"
	"strconv"
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

	return context.JSON(common.NewSuccessResponseWithoutData(common.SUCCESS_INSERT_DATA))
}

func (controller *Controller) GetListLessonType(context echo.Context) (err error) {
	var listLessonType []response.GetListLessonType
	listLessonType, err = controller.service.GetListLessonType()
	if err != nil {
		return
	}

	return context.JSON(common.NewSuccessResponse(common.SUCCESS_GET_LIST_DATA, listLessonType))
}

func (controller *Controller) UpdateLessonType(context echo.Context) error {
	var request request.UpdateLessonType

	// binding ke request dan validation
	if err := context.Bind(&request); err != nil {
		return context.JSON(common.NewErrBindData())
	}

	dataFailed, isFailed := request.MandatoryValidation()
	if isFailed {
		return context.JSON(common.NewValidationResponse(dataFailed))
	}

	// mengambil id dari url param
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}
	request.ID = id

	err = controller.service.UpdateLessonType(request)
	if err != nil {
		return context.JSON(common.NewErrorBusinessResponse(err))
	}

	return context.JSON(common.NewSuccessResponseWithoutData(common.SUCCESS_UPDATE_DATA))
}