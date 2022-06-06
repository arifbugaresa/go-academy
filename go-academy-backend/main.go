package main

import (
	configuration "go-academy/config"
	"go-academy/modules/database"
)

func main()  {
	// load configuration
	config := configuration.GetConfig()

	// initialize database connection
	dbGormPostgres := database.NewDatabaseConnection(config)
	defer database.CloseDatabaseConnection(dbGormPostgres)

}