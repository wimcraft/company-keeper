package user

import "company-keeper-go/models"

type Store interface {
	GetByID(uint) (*models.User, error)
	Create(*models.User) error
	GetByEmail(string) (*models.User, error)
	Delete(*models.User) error
}
