package models

import (
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Nombre          string `json:"nombre"`
	Paterno         string `json:"paterno"`
	Materno         string `json:"materno"`
	Usuario         string `json:"usuario"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Edad            string `json:"edad"`
	Fechanacimiento string `json:"fechanacimiento"`
	DNI             string `json:"dni"`
}
