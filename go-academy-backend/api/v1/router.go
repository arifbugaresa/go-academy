package api

import (
	"github.com/labstack/echo/v4"
	"go-academy/api/v1/lesson_group"
	"go-academy/api/v1/lesson_type"
)

func Controller(
	e *echo.Echo,
	lessonTypeController *lesson_type.Controller,
	lessonGroupController *lesson_group.Controller,
) {

	// helper
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

	// lesson type
	lessonType := e.Group("v1/lesson-type")
	lessonType.POST("", lessonTypeController.InsertLessonType)
	lessonType.GET("", lessonTypeController.GetListLessonType)
	lessonType.PUT("/:id", lessonTypeController.UpdateLessonType)
	lessonType.DELETE("/:id", lessonTypeController.DeleteLessonTypeByID)

	// lesson group
	lessonGroup := e.Group("v1/lesson-group")
	lessonGroup.POST("", lessonGroupController.InsertLessonGroup)
	lessonGroup.GET("", lessonGroupController.FindAllLessonGroup)

}
