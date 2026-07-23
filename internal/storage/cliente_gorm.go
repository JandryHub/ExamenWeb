package storage

import (
	"gorm.io/gorm"

	"github.com/joancema/examen-floristeria/internal/models"
)

// TAREA (CP1): Implemente ClienteGORM contra la interfaz ClienteRepository.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos:
//     los tests de acceptance/ compilan contra ellos.
//   - Guíese por ArregloGORM (arreglo_gorm.go): es el mismo patrón.
type ClienteGORM struct {
	db *gorm.DB
}

func NuevoClienteGORM(db *gorm.DB) *ClienteGORM {
	return &ClienteGORM{db: db}
}

func (r *ClienteGORM) Crear(c *models.Cliente) error {
	return r.db.Create(c).Error
}

func (r *ClienteGORM) ObtenerPorID(id uint) (models.Cliente, bool) {
	var c models.Cliente
	result := r.db.First(&c, id)
	return c, result.Error == nil
}

func (r *ClienteGORM) Listar() ([]models.Cliente, error) {
	var clientes []models.Cliente
	err := r.db.Find(&clientes).Error
	return clientes, err
}
