package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go-academy/api/v1"
	"go-academy/business/lesson"
	"go-academy/business/lesson_group"
	"go-academy/business/lesson_type"
	configuration "go-academy/config"
	"go-academy/modules/database"
	"strconv"
)

var (
	config         = configuration.GetConfig()
	dbGormPostgres = database.NewDatabaseConnection(config)
	e              = echo.New()
)

func main() {
	defer database.CloseDatabaseConnection(dbGormPostgres)

	migrateDatabase()

	api.Controller(e)

	runServer()
}

func migrateDatabase() {
	dbGormPostgres.AutoMigrate(
		&lesson.Lesson{},
		&lesson_type.LessonType{},
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