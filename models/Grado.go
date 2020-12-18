package models

import "gorm.io/gorm"

type Grado struct {
	gorm.Model
	Nombre  string `json:"nombre"`
	NivelID string `gorm:"size:191"`
	Nivel   Nivel
	Estado  string `json:"estado"`
}
