package handler

import (
	"company-keeper-go/router/middleware"
	"company-keeper-go/utils"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	guestUsers := v1.Group("/users")
	guestUsers.POST("", h.SignUp)
	guestUsers.POST("/login", h.Login)

	jwtMiddleware := middleware.JWTWithConfig(
		middleware.JWTConfig{
			SigningKey: utils.GetJWTSecret(),
		},
	)
	user := v1.Group("/user", jwtMiddleware)
	user.GET("", h.CurrentUser)

	companies := v1.Group("/company",
		middleware.JWTWithConfig(
			middleware.JWTConfig{
				Skipper: func(c echo.Context) bool {
					return c.Request().Method == "GET"
				},
				SigningKey: utils.GetJWTSecret(),
			},
		))
	companies.GET("/:companyId", h.CompanyInfo)
	companies.POST("", h.CreateCompany)
	companies.PATCH("/:companyId", h.UpdateCompany)
	companies.DELETE("/:companyId", h.DeleteCompany)

}
