package sqlapp

import (
	"database/sql"
	"web/iternal/app"
)

type App struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *App {
	return &App{db: db}
}

func (a *App) User() app.UserRepository {
	if a.userRepository != nil {
		return a.userRepository
	}
	a.userRepository = &UserRepository{
		app: a,
	}
	return a.userRepository
}
