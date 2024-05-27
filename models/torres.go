package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Torre struct {
	gorm.Model
	Condominio_id              uint   `json:"condominio_id"`
	Numero_torre               int    `json:"numero_torre"`
	Nome_torre                 string `json:"nome_torre"`
	Quantidade_andares         int    `json:"quantidade_andares"`
	Quantidade_garagens        int    `json:"quantidade_garagens"`
	Quantidade_salao_de_festas int    `json:"quantidade_salao_de_festas"`
	Quantidade_guaritas        int    `json:"quantidade_guaritas"`
	Quantidade_terracos        int    `json:"quantidade_terracos"`
}

func ValidatesDataTorres(torre *Torre) error {
	if err := validator.Validate(torre); err != nil {
		return err
	}
	return nil
}
