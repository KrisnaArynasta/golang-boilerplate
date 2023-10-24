package main

import (
	"databaseservice"
	"time"

	"github.com/gin-gonic/gin"

	Route "testing-api/api/route"
	Bootstrap "testing-api/bootstrap"
	Loghelper "testing-api/loghelper"
)

func main() {
	Loghelper.WriteLog().Info().Msg("Application Start")

	app := Bootstrap.App()
	env := app.Env

	// db := app.SqlDb
	// defer app.CloseDBConnection()
	db := databaseservice.Init(env.DBHost, env.DBUser, env.DBPass, env.DBName, env.DBPort)
	defer db.CloseDBConnection()

	gin := gin.Default()
	timeout := time.Duration(2) * time.Second

	Route.Setup(gin, timeout, db, env)
	gin.Run()
}
