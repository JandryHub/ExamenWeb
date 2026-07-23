// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import "github.com/joancema/examen-floristeria/internal/models"

// ArregloRepository define el contrato de persistencia de la Entidad A.
type ArregloRepository interface {
	Crear(h *models.Arreglo) error
	ObtenerPorID(id uint) (models.Arreglo, bool)
	Listar() ([]models.Arreglo, error)
	Actualizar(h *models.Arreglo) error
}
