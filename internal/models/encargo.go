package models

import "gorm.io/gorm"

// TAREA (CP1): Complete los campos de Encargo según lo que muestran las pantallas.
//
// Pistas de trabajo:
//   - Un Encargo referencia a un Arreglo y a un Cliente (claves foráneas).
//   - Recuerde el campo de estado (use las constantes de estados.go) y el total.
//   - Los tests de acceptance/ compilan contra los nombres EXACTOS de los campos.
type Encargo struct {
	gorm.Model
	ArregloID uint    `gorm:"not null" json:"arreglo_id"`
	Arreglo   Arreglo `json:"arreglo"`
	ClienteID uint    `gorm:"not null" json:"cliente_id"`
	Cliente   Cliente `json:"cliente"`
	Cantidad  uint    `gorm:"not null" json:"cantidad"`
	Estado    string  `gorm:"not null" json:"estado"`
	Total     float64 `gorm:"not null" json:"total"`
}
