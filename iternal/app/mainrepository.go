package app

import "web/models"

type UserRepository interface {
	Create(user *models.User) error
}
