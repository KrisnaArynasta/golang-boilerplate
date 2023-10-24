package Bootstrap

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	Loghelper "testing-api/loghelper"

// 	_ "github.com/denisenkom/go-mssqldb"
// )

// func NewDatabaseMsSql(env *Env) *sql.DB {
// 	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	// defer cancel()

// 	dbHost := env.DBHost
// 	dbUser := env.DBUser
// 	dbPass := env.DBPass
// 	dbName := env.DBName
// 	dbPort := env.DBPort

// 	connectionStr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
// 		dbHost, dbUser, dbPass, dbPort, dbName)

// 	db, err := sql.Open("mssql", connectionStr)
// 	if err != nil {
// 		fmt.Println("Error connecting to database:", err)
// 		Loghelper.WriteLog().Error().Msg(fmt.Sprintf("Error connecting to database: %s", err))
// 	}

// 	return db
// }

// func CloseMsSqlConnection(client *sql.DB) {
// 	if client == nil {
// 		return
// 	}

// 	err := client.Close()
// 	if err != nil {
// 		Loghelper.WriteLog().Error().Msg(err.Error())
// 	}

// 	Loghelper.WriteLog().Info().Msg("Connection to Ms SQL closed")
// }

/* ======================================================================================*/
/* ===================== MOVED TO OTHER PROJECT (DATABASE-SERIVCE) =====================*/
/* ====================================================================================*/
