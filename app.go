package app

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize takes in the details required to connect to the database.
// It will create a database connection and wire up the routes to
// respond to the requirements.
func (a *App) Initialize(user, password, dbname string) {}

// Run simply starts the application.
func (a *App) Run(address string) {}
