package handler

import (
	"company-keeper-go/models"
	"company-keeper-go/utils"
	"github.com/google/uuid"
)

type userResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func newUserResponse(u *models.User) *userResponse {
	r := new(userResponse)
	r.Username = u.Username
	r.Email = u.Email
	r.Token = utils.GenerateJWT(u.ID)
	return r
}

type companyResponse struct {
	ID           uuid.UUID `json:"id"`
	Description  string    `json:"description"`
	Name         string    `json:"name"`
	NumEmployees int       `json:"numEmployees"`
	Registered   bool      `json:"registered"`
	Type         string    `json:"type"`
}

func newCompanyResponse(u *models.Company) *companyResponse {
	r := new(companyResponse)
	r.ID = u.ID
	r.Description = u.Description
	r.Name = u.Name
	r.NumEmployees = u.NumEmployees
	r.Registered = u.Registered
	r.Type = u.Type
	return r
}
