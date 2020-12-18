package models

import (
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Nombre          string `json:"nombre"`
	Paterno         string `json:"paterno"`
	Materno         string `json:"materno"`
	Email           string `json:"email"`
	Edad            string `json:"edad"`
	Celular         string `json:"celular"`
	Fechanacimiento string `json:"fechanacimiento"`
	DNI             string `json:"dni"`
	Estado          string `json:"estado"`
}
