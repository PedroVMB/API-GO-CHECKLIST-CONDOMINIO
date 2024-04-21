package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Sindico struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
}

func ValidatesDataSindico(sindico *Sindico) error {
	if err := validator.Validate(sindico); err != nil {
		return err
	}
	return nil
}
