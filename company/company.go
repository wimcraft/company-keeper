package company

import (
	"company-keeper-go/models"
	"github.com/google/uuid"
)

type Store interface {
	GetByID(uuid uuid.UUID) (*models.Company, error)
	Create(company *models.Company) error
	Update(company *models.Company) error
	Delete(company *models.Company) error
}
