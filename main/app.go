package main

import (
	"database/sql"
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
	// var connectionString string

	// if len(password) > 0 {
	// 	connectionString = fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", user, password, dbname)
	// } else {
	// 	connectionString = fmt.Sprintf("postgres://%s@localhost/%s?sslmode=disable", user, dbname)
	// }

	var err error
	a.DB, err = sql.Open("postgres", "postgres://jimxshaw@localhost/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

// Run simply starts the application.
func (a *App) Run(address string) {}
