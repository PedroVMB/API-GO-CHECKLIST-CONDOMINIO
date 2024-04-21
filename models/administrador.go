package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Administrador struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
}

func ValidatesDataAdm(adm *Administrador) error {
	if err := validator.Validate(adm); err != nil {
		return err
	}
	return nil
}
