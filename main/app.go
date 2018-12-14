package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"

	// pq is here because we need our
	// application to work with PostgreSQL.
	_ "github.com/lib/pq"
)

// App type represents what this
// application is.
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize takes in the details required to connect to the database.
// It will create a database connection and wire up the routes to
// respond to the requirements.
func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

// Run simply starts the application.
func (a *App) Run(address string) {}
