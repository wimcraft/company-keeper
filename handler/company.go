package handler

import (
	"company-keeper-go/models"
	"company-keeper-go/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CreateCompany godoc
// @Summary Register a new company
// @Description Register a new company
// @ID create-company
// @Tags company
// @Accept  json
// @Produce  json
// @Param company body companyRegisterRequest true "Company info for registration"
// @Success 201 {object} companyResponse
// @Failure 400 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /company [post]
func (h *Handler) CreateCompany(c echo.Context) error {
	var company models.Company
	req := &companyRegisterRequest{}
	if err := req.bind(c, &company); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := h.companyStore.Create(&company); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newCompanyResponse(&company))
}

// UpdateCompany godoc
// @Summary Update a new company
// @Description Register a new company
// @ID update-company
// @Tags company
// @Accept  json
// @Produce  json
// @Param company body companyUpdateRequest true "Company data for update"
// @Success 201 {object} companyResponse
// @Failure 400 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /company [patch]
func (h *Handler) UpdateCompany(c echo.Context) error {
	companyId := c.Param("companyId")

	companyIdConverted, err := uuid.Parse(companyId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	var company *models.Company
	company, err = h.companyStore.GetByID(companyIdConverted)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if company == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	req := &companyUpdateRequest{ID: companyIdConverted}
	req.populate(company)

	if err := req.bind(c, company); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if err = h.companyStore.Update(company); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, newCompanyResponse(company))
}

// CompanyInfo godoc
// @Summary Get company info
// @Description Get company info
// @ID info-company
// @Tags company
// @Accept  json
// @Produce  json
// @Param companyId  path string true "Company ID"
// @Success 200 {object} userResponse
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /company/{companyId} [get]
func (h *Handler) CompanyInfo(c echo.Context) error {
	companyId := c.Param("companyId")

	companyIdConverted, err := uuid.Parse(companyId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	var company *models.Company
	company, err = h.companyStore.GetByID(companyIdConverted)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if company == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	return c.JSON(http.StatusOK, newCompanyResponse(company))
}

// DeleteCompany godoc
// @Summary Delete company
// @Description Delete company
// @ID delete-company
// @Tags company
// @Accept  json
// @Produce  json
// @Param companyId path string true "Slug of the article to delete"
// @Success 201 {object} companyResponse
// @Failure 400 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /company/{companyId} [delete]
func (h *Handler) DeleteCompany(c echo.Context) error {
	companyId := c.Param("companyId")

	companyIdConverted, err := uuid.Parse(companyId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	var company *models.Company
	company, err = h.companyStore.GetByID(companyIdConverted)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if company == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	err = h.companyStore.Delete(company)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}
