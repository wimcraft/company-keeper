package store

import (
	"company-keeper-go/models"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyStore struct {
	db *gorm.DB
}

func NewCompanyStore(db *gorm.DB) *CompanyStore {
	return &CompanyStore{
		db: db,
	}
}

func (co *CompanyStore) GetByID(id uuid.UUID) (*models.Company, error) {
	var m models.Company
	if err := co.db.First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (co *CompanyStore) Create(u *models.Company) (err error) {
	return co.db.Create(u).Error
}

func (co *CompanyStore) Update(u *models.Company) error {
	// Save the updated company to the database
	err := co.db.Save(u).Error
	return err
}

func (co *CompanyStore) Delete(u *models.Company) error {
	return co.db.Model(u).Delete(u).Error
}
