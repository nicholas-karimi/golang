package config

import (
	"log"
	"os"
)

// app struct to hold appliction wide dependencies
type Application struct {
	errLog  *log.Logger
	infoLog *log.Logger
}

func (app *Application) LogInfo(message string) {
	app.infoLog.Printf(message)

}

func (app *Application) LogError(message string) {
	app.errLog.Printf(message)

}

// NewLogger creates a new instance of CustomLogger.
func NewLogger() *Application {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return &Application{
		infoLog: infoLog,
		errLog:  errLog,
	}
}
