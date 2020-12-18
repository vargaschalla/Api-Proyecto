package models

import "gorm.io/gorm"

type Nivel struct {
	gorm.Model
	Nombre string `json:"nombre"`
	Grado  []Grado
	Estado string `json:"estado"`
}
