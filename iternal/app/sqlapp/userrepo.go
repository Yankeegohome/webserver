package sqlapp

import "web/models"

type UserRepository struct {
	app *App
}

func (u *UserRepository) Create(mu *models.User) error {
	return u.app.db.QueryRow(
		"INSERT INTO USERS(ID, EMAIL, PASSWORD) VALUES (1, $1, $2) RETURNING id",
		mu.Login,
		mu.Password,
	).Scan(&mu.ID)
}
