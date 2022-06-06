package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go-academy/api/v1"
	lessonTypeContr "go-academy/api/v1/lesson_type"
	"go-academy/business/lesson"
	"go-academy/business/lesson_group"
	lessonType "go-academy/business/lesson_type"
	lessonTypeServ "go-academy/business/lesson_type"
	configuration "go-academy/config"
	"go-academy/modules/database"
	lessonTypeRepo "go-academy/modules/lesson_type"
	"strconv"
)

var (
	config         = configuration.GetConfig()
	dbGormPostgres = database.NewDatabaseConnection(config)
	e              = echo.New()
)

var (
	lessonTypeRepository = lessonTypeRepo.NewRepository(dbGormPostgres)
	lessonTypeService    = lessonTypeServ.NewService(lessonTypeRepository)
	lessonTypeController = lessonTypeContr.NewController(lessonTypeService)
)

func main() {
	defer database.CloseDatabaseConnection(dbGormPostgres)

	migrateDatabase()

	api.Controller(e, lessonTypeController)

	runServer()
}

func migrateDatabase() {
	dbGormPostgres.AutoMigrate(
		&lesson.Lesson{},
		&lessonType.LessonType{},
		&lesson_group.LessonGroup{},
	)

	log.Info("Success migrate database, " + strconv.Itoa(int(dbGormPostgres.RowsAffected)) + " row affected.")
}

func runServer() {
	address := fmt.Sprintf("localhost:%s", config.AppPort)
	err := e.Start(address)
	if err != nil {
		log.Info("shutting down the server")
	}
}
