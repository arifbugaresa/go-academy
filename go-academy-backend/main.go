package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go-academy/api/v1"
	lessonContr "go-academy/api/v1/lesson"
	lessonGroupContr "go-academy/api/v1/lesson_group"
	lessonTypeContr "go-academy/api/v1/lesson_type"
	lessonServ "go-academy/business/lesson"
	lessonGroupServ "go-academy/business/lesson_group"
	lessonTypeServ "go-academy/business/lesson_type"
	configuration "go-academy/config"
	"go-academy/modules/database"
	lessonRepo "go-academy/modules/lesson"
	lessonGroupRepo "go-academy/modules/lesson_group"
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

	lessonGroupRepository = lessonGroupRepo.NewRepository(dbGormPostgres)
	lessonGroupService    = lessonGroupServ.NewService(lessonGroupRepository)
	lessonGroupController = lessonGroupContr.NewController(lessonGroupService)

	lessonRepository = lessonRepo.NewRepository(dbGormPostgres)
	lessonService    = lessonServ.NewService(lessonRepository)
	lessonController = lessonContr.NewController(lessonService)
)

func main() {
	defer database.CloseDatabaseConnection(dbGormPostgres)

	migrateDatabase()

	api.Controller(e, lessonTypeController, lessonGroupController, lessonController)

	runServer()
}

func migrateDatabase() {
	dbGormPostgres.AutoMigrate(
		&lessonServ.Lesson{},
		&lessonTypeServ.LessonType{},
		&lessonGroupServ.LessonGroup{},
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
