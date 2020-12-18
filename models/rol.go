package models

import (
	"gorm.io/gorm"
)

type Rol struct {
	gorm.Model
	Nombre  string `json:"nombre"`
	Alumno  []Alumno
	Docente []Docente
	Estado  string `json:"estado"`
}
