package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nicholas-karimi/snippetbox/cmd/web/handlers"
	"github.com/nicholas-karimi/snippetbox/config"
	

	"github.com/nicholas-karimi/snippetbox/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// logger instance
	logger := config.NewLogger()

	// command line flag - default value :4000 & help  text to expalin what the flag controls
	addr := flag.String("addr", ":4000", "HTTP network address")

	// dsn := flag.String("dsn", "admin:Incorrect@/snippetbox?parseTime=true", "MYSQL Source name to connect...")

	// db, err := openDB(*dsn)
	// if err != nil {
	// 	logger.LogError(err.Error())
	// 	os.Exit(1)
	// }

	// defer db.Close()

	// parse the commabndline flag
	flag.Parse()

	// Initialize database connection
	db := config.Db

	// Initialize SnippetModel
    snippetModel := &models.SnippetModel{DB: db}

	// Initialize AppConfig
	appConfig := &config.AppConfig{
		InfoLog:      logger.InfoLog,
		ErrorLog:     logger.ErrorLog,
		SnippetModel: snippetModel, 
	}

	// Initialize Handlers with AppConfig
	appHandlers := &handlers.Handlers{AppConfig: appConfig}

	// use log.New() to create a logger for writing information messages.
	// takes 3 params: destination to write logs(os.Stdout), string prefix for message(INFO)
	//and tags for additional ifno to include (local date and time - flags joined using bitwise OR operator |)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// error msg logger - use stderr as destination and log.Lshortfile flag to include relevant filebame and line number
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)



	mux := http.NewServeMux()

	// file server to serve files from the static directory
	fileServer := http.FileServer(http.Dir("ui/static"))

	// register the file server as the handler for the static directory
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// other app routes
	/* mux.HandleFunc("/", handlers.Home(logger))
	mux.HandleFunc("/snippet/view", handlers.SnippetView)
	mux.HandleFunc("/snippet/create", handlers.SnippetCreate) */
	mux.HandleFunc("/", appHandlers.Home)
	mux.HandleFunc("/snippet/view", appHandlers.SnippetView)
	mux.HandleFunc("/snippet/create", appHandlers.SnippetCreate)

	/* // log.Print("Starting web server on :4000")
	// log.Printf("Starting server on %s", *addr)
	infoLog.Printf("Starting server on %s", *addr)

	// err := http.ListenAndServe(":4000", mux)
	err := http.ListenAndServe(*addr, mux)

	// log.Fatal(err) */
	// Print the value of AppConfig and SnippetModel for debugging
    fmt.Println("AppConfig:", appConfig)
    fmt.Println("SnippetModel:", appConfig.SnippetModel)

	// http.Server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  mux,
	}
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()

	errLog.Fatal(err)

	// db, err := config.InitDB()
	// if err != nil {
	// 	logger.LogError(err.Error())
	// 	os.Exit(1)
	// }

	// defer db.Close()
	fmt.Println("Connected to MySQL database!")
}
