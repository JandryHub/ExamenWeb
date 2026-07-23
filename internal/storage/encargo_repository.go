// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import "github.com/joancema/examen-floristeria/internal/models"

// EncargoRepository define el contrato de persistencia de Encargo.
// Su implementación GORM (en encargo_gorm.go) debe satisfacer EXACTAMENTE
// estas firmas. Observe que el repositorio NO contiene lógica de negocio:
// las reglas (validaciones, cálculo del total, anulación) viven en el service.
type EncargoRepository interface {
	Crear(a *models.Encargo) error
	ObtenerPorID(id uint) (models.Encargo, bool)
	Listar() ([]models.Encargo, error)
	Actualizar(a *models.Encargo) error
}
