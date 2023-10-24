package Bootstrap

import "database/sql"

type Application struct {
	Env   *Env
	SqlDb *sql.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	//app.SqlDb = NewDatabaseMsSql(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	//CloseMsSqlConnection(app.SqlDb)
}
