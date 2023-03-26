package handler

import (
	"company-keeper-go/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type userRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *userRegisterRequest) bind(c echo.Context, u *models.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Email = r.Email
	u.Username = r.Username
	h, err := u.HashPassword(r.Password)
	if err != nil {
		return err
	}
	u.Password = h
	return nil
}

type userLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *userLoginRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

type companyRegisterRequest struct {
	Description  string `json:"description"`
	Name         string `json:"name"`
	NumEmployees int    `json:"numEmployees" validate:"required"`
	Registered   bool   `json:"registered" validate:"required"`
	Type         string `json:"type" validate:"required,companytype"`
}

func (r *companyRegisterRequest) bind(c echo.Context, u *models.Company) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	u.Description = r.Description
	u.Name = r.Name
	u.NumEmployees = r.NumEmployees
	u.Registered = r.Registered
	u.Type = r.Type
	return nil
}

type companyUpdateRequest struct {
	ID           uuid.UUID `json:"id"`
	Description  string    `json:"description"`
	Name         string    `json:"name"`
	NumEmployees int       `json:"numEmployees" validate:"required"`
	Registered   bool      `json:"registered" validate:"required"`
	Type         string    `json:"type" validate:"required,companytype"`
}

func (r *companyUpdateRequest) populate(co *models.Company) {
	r.Description = co.Description
	r.Name = co.Name
	r.NumEmployees = co.NumEmployees
	r.Registered = co.Registered
	r.Type = co.Type
}

func (r *companyUpdateRequest) bind(c echo.Context, company *models.Company) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	company.Description = r.Description
	company.Name = r.Name
	company.NumEmployees = r.NumEmployees
	company.Registered = r.Registered
	company.Type = r.Type

	return nil
}
