package Route

import (
	"databaseservice"
	Controller "testing-api/api/controller"
	Bootstrap "testing-api/bootstrap"
	Database "testing-api/database"
	Repository "testing-api/repository"
	Service "testing-api/service"
	"time"

	"github.com/gin-gonic/gin"
)

func TransactionRoute(group *gin.RouterGroup, protectedRouter *gin.RouterGroup, timeout time.Duration, db *databaseservice.DatabaseService, env *Bootstrap.Env) {
	tr := Repository.NewTransactionRepository(env)
	td := Database.NewTransactionDatabase(db)
	tc := &Controller.TransactionController{
		TransactionService: Service.NewTransactionService(tr, timeout),
		DataBase:           td,
	}

	group.GET("/get_data", tc.LoadData)

	// protectedRouter will use Middleware.BasicAuth (in route.Setup) for private request validation.
	protectedRouter.POST("/post_data", tc.PostData)
	protectedRouter.GET("/get_db", tc.LoadDataFromDatabase)
}
