// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import (
	"gorm.io/gorm"

	"github.com/joancema/examen-floristeria/internal/models"
)

// ArregloGORM implementa ArregloRepository sobre GORM.
// Esta implementación está completa: úsela como plantilla para ClienteGORM
// y EncargoGORM, que usted debe implementar.
type ArregloGORM struct {
	db *gorm.DB
}

func NuevoArregloGORM(db *gorm.DB) *ArregloGORM {
	return &ArregloGORM{db: db}
}

func (r *ArregloGORM) Crear(h *models.Arreglo) error {
	return r.db.Create(h).Error
}

func (r *ArregloGORM) ObtenerPorID(id uint) (models.Arreglo, bool) {
	var h models.Arreglo
	if err := r.db.First(&h, id).Error; err != nil {
		return models.Arreglo{}, false
	}
	return h, true
}

func (r *ArregloGORM) Listar() ([]models.Arreglo, error) {
	var lista []models.Arreglo
	err := r.db.Find(&lista).Error
	return lista, err
}

func (r *ArregloGORM) Actualizar(h *models.Arreglo) error {
	return r.db.Save(h).Error
}
