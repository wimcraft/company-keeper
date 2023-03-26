package router

import "gopkg.in/go-playground/validator.v9"

func ValidateCompanyType(fl validator.FieldLevel) bool {
	companyType := fl.Field().String()
	allowedTypes := []string{"Corporations", "NonProfit", "Cooperative", "Sole Proprietorship"}
	for _, t := range allowedTypes {
		if t == companyType {
			return true
		}
	}
	return false
}

func NewValidator() *Validator {
	validate := validator.New()
	err := validate.RegisterValidation("companytype", ValidateCompanyType)
	if err != nil {
		return nil
	}
	return &Validator{
		validator: validate,
	}
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
