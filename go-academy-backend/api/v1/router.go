package api

import (
	"github.com/labstack/echo/v4"
	"go-academy/api/v1/lesson_type"
)

func Controller(
	e *echo.Echo,
	lessonTypeControler *lesson_type.Controller,
) {

	// helper
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

	// lesson type
	lessonType := e.Group("v1/lesson-type")
	lessonType.POST("", lessonTypeControler.InsertLessonType)
	lessonType.GET("", lessonTypeControler.GetListLessonType)
	lessonType.PUT("/:id", lessonTypeControler.UpdateLessonType)
	lessonType.DELETE("/:id", lessonTypeControler.DeleteLessonTypeByID)

}
