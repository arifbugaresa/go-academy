package main

import (
	"github.com/labstack/gommon/log"
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
)

func main() {
	defer database.CloseDatabaseConnection(dbGormPostgres)

	migrate()

}

func migrate() {
	dbGormPostgres.AutoMigrate(
		&lesson.Lesson{},
		&lesson_type.LessonType{},
		&lesson_group.LessonGroup{},
	)

	log.Info("Success migrate database, " + strconv.Itoa(int(dbGormPostgres.RowsAffected)) + " row affected.")
}