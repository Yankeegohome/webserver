package webserver

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log/slog"
	"net/http"
	"web/iternal/app/sqlapp"
)

func Start(config *Config) error {
	db, err := NewDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	app := sqlapp.New(db)
	srv := newServer(app)
	return http.ListenAndServe(config.BindAddr, srv)
}

func NewDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", databaseURL)
	slog.Info("trying  connect to db")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	slog.Info("db connected: " + databaseURL)

	return db, nil
}
