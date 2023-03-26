package handler

import (
	"company-keeper-go/models"
	"company-keeper-go/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCompanyCaseSuccess(t *testing.T) {
	setup()
	var (
		reqJSON = `{"description": "Test description","name": "Acme", "numEmployees": 5, "registered": true, "type": "NonProfit"}`
	)
	req := httptest.NewRequest(echo.POST, "/api/company", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(1)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, h.CreateCompany(c))
	if assert.Equal(t, http.StatusCreated, rec.Code) {
		m := responseMap(rec.Body.Bytes())
		assert.NotEmpty(t, m["id"])
		assert.Equal(t, "Acme", m["name"])
		assert.Equal(t, "Test description", m["description"])
		assert.Equal(t, float64(5), m["numEmployees"])
		assert.Equal(t, true, m["registered"])
		assert.Equal(t, "NonProfit", m["type"])
	}
}

func TestUpdateCompanyCaseSuccess(t *testing.T) {
	setup()
	var (
		reqJSON = `{"numEmployees": 7}`
	)

	company := models.Company{
		Name:         "TestCompany",
		Description:  "No description",
		NumEmployees: 10,
		Registered:   true,
		Type:         "NonProfit",
	}

	if err := co.Create(&company); err != nil {
		t.Error(err)
	}

	req := httptest.NewRequest(echo.PATCH, "/api/company/:companyId", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(1)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/company/:companyId")
	c.SetParamNames("companyId")
	c.SetParamValues(company.ID.String())
	assert.NoError(t, h.UpdateCompany(c))
	if assert.Equal(t, http.StatusOK, rec.Code) {
		m := responseMap(rec.Body.Bytes())
		assert.Equal(t, company.ID.String(), m["id"])
		assert.Equal(t, "TestCompany", m["name"])
		assert.Equal(t, "No description", m["description"])
		assert.Equal(t, float64(7), m["numEmployees"])
		assert.Equal(t, true, m["registered"])
		assert.Equal(t, "NonProfit", m["type"])
	}
}

func TestGetCompanyCaseSuccess(t *testing.T) {
	setup()
	company := models.Company{
		Name:         "TestCompany",
		Description:  "No description",
		NumEmployees: 10,
		Registered:   true,
		Type:         "NonProfit",
	}

	if err := co.Create(&company); err != nil {
		t.Error(err)
	}

	req := httptest.NewRequest(echo.GET, "/api/company/:companyId", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/company/:companyId")
	c.SetParamNames("companyId")
	c.SetParamValues(company.ID.String())
	assert.NoError(t, h.CompanyInfo(c))
	if assert.Equal(t, http.StatusOK, rec.Code) {
		m := responseMap(rec.Body.Bytes())
		assert.Equal(t, company.ID.String(), m["id"])
		assert.Equal(t, "TestCompany", m["name"])
		assert.Equal(t, "No description", m["description"])
		assert.Equal(t, float64(10), m["numEmployees"])
		assert.Equal(t, true, m["registered"])
		assert.Equal(t, "NonProfit", m["type"])
	}
}

func TestDeleteCompanyCaseSuccess(t *testing.T) {
	setup()
	company := models.Company{
		Name:         "TestCompany",
		Description:  "No description",
		NumEmployees: 10,
		Registered:   true,
		Type:         "NonProfit",
	}

	if err := co.Create(&company); err != nil {
		t.Error(err)
	}

	req := httptest.NewRequest(echo.DELETE, "/api/company/:companyId", nil)
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(1)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/company/:companyId")
	c.SetParamNames("companyId")
	c.SetParamValues(company.ID.String())
	assert.NoError(t, h.DeleteCompany(c))
	if assert.Equal(t, http.StatusOK, rec.Code) {
		m := responseMap(rec.Body.Bytes())
		assert.Equal(t, "ok", m["result"])
	}
}
