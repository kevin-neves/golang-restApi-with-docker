package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
	Cpf  string `json:"cpf"`
}
