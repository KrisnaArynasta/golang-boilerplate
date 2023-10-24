package Route

import (
	"databaseservice"
	Bootstrap "testing-api/bootstrap"
	Middleware "testing-api/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(gin *gin.Engine, timeout time.Duration, db *databaseservice.DatabaseService, env *Bootstrap.Env) {
	publicRouter := gin.Group("")

	protectedRouter := gin.Group("")
	protectedRouter.Use(Middleware.BasicAuth(env.AccessTokenSecret))

	TransactionRoute(publicRouter, protectedRouter, timeout, db, env)
}
