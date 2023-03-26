package handler

import (
	"company-keeper-go/company"
	"company-keeper-go/user"
)

type Handler struct {
	userStore    user.Store
	companyStore company.Store
}

func NewHandler(us user.Store, co company.Store) *Handler {
	return &Handler{
		userStore:    us,
		companyStore: co,
	}
}
