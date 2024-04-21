package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Condominio struct {
	gorm.Model
	Nome              string `json:"name"`
	CNPJ              string `json:"cnpj" validate:"required,len=14,regex=^\d{14}$"`
	Estado            string `json:"estado"`
	Cidade            string `json:"cidade"`
	Bairro            string `json:"bairro"`
	Numero            int    `json:"numero"`
	CEP               string `json:"cep" validate:"required,len=8,regex=^\d{8}$"`
	Quantidade_torres int    `json:"quantidade_torres"`
}

func ValidatesDataCondominio(condominio *Condominio) error {
	if err := validator.Validate(condominio); err != nil {
		return err
	}
	return nil
}
