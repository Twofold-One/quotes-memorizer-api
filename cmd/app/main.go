package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Twofold-One/quotes-memorizer-api/pkg/models/postgresql"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type App struct {
	DB *sql.DB
	Router *mux.Router
	quotes *postgresql.QuoteModel
}

func (app *App) InitializeApp() {
	db, err := sql.Open("pgx", os.Getenv("DBURL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to DB!")

	app.Router = mux.NewRouter()
	app.InitializeRoutes()
}

func (app *App) Run() {
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), app.Router))
}

func main() {

	app := App{}
	app.InitializeApp()
	app.Run()

	// custom loggers
	// infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	
}

