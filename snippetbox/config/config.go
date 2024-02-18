package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/nicholas-karimi/snippetbox/internal/models"

	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/lib/pq"
)

// database configuration - get package level Db variable- avoild nil point
var Db *sql.DB

// func InitDB() (*sql.DB, error) {
// 	// dsn := "user=admin dbname=snippetbox password=Incorrect host=localhost port=5432 sslmode=disable"
// 	// dsn := "admin:Incorrect@tcp(127.0.0.1:3306)/snippetbox?charset=utf8mb4&parseTime=True&loc=Local"
// 	// Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	dsn := "admin:Incorrect@tcp(localhost:3306)/snippetbox?parseTime=true"
// 	db, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := db.Ping(); err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }

func init() {
	dsn := "admin:Incorrect@tcp(localhost:3306)/snippetbox?parseTime=true"
	var err error
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}

}

// app struct to hold appliction wide dependencies
// type Application struct {
// 	errLog   *log.Logger
// 	infoLog  *log.Logger
// 	snippets *models.SnippetModel
// }

// func (app *Application) LogInfo(message string) {
// 	app.infoLog.Printf(message)

// }

// func (app *Application) LogError(message string) {
// 	app.errLog.Printf(message)

// }

// func (app *Application) clientError(w http.ResponseWriter, status int) {
// 	http.Error(w, http.StatusText(status), status)
// }

// func (app *Application) notFound(w http.ResponseWriter) {
// 	app.clientError(w, http.StatusNotFound)
// }

// // NewLogger creates a new instance of CustomLogger.
// func NewLogger() *Application {
// 	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
// 	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

// 	return &Application{
// 		infoLog:  infoLog,
// 		errLog:   errLog,
// 		// snippets: &models.SnippetModel{DB: Db},
// 	}
// }

// CustomLogger represents application-wide loggers
type CustomLogger struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

// NewLogger initializes and returns a new instance of CustomLogger
func NewLogger() *CustomLogger {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return &CustomLogger{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}
}

// AppConfig holds application configurations and dependencies
type AppConfig struct {
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	SnippetModel *models.SnippetModel
}

// InitializeAppConfig initializes and returns a new instance of AppConfig
func InitializeAppConfig() *AppConfig {
	return &AppConfig{
		InfoLog:      NewLogger().InfoLog,
		ErrorLog:     NewLogger().ErrorLog,
		SnippetModel: &models.SnippetModel{DB: Db},
	}
}
