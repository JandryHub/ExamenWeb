package storage

import (
	"errors"

	"gorm.io/gorm"

	"github.com/joancema/examen-floristeria/internal/models"
)

// TAREA (CP2): Implemente EncargoGORM contra la interfaz EncargoRepository.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos.
//   - Guíese por ArregloGORM: es el mismo patrón con una entidad distinta.
//   - Recuerde: aquí NO va lógica de negocio. Solo persistencia.
type EncargoGORM struct {
	db *gorm.DB
}

func NuevoEncargoGORM(db *gorm.DB) *EncargoGORM {
	return &EncargoGORM{db: db}
}

func (r *EncargoGORM) Crear(a *models.Encargo) error {
	// TODO: implementar.
	return errors.New("TODO: implementar Crear")
}

func (r *EncargoGORM) ObtenerPorID(id uint) (models.Encargo, bool) {
	// TODO: implementar.
	return models.Encargo{}, false
}

func (r *EncargoGORM) Listar() ([]models.Encargo, error) {
	// TODO: implementar.
	return nil, errors.New("TODO: implementar Listar")
}

func (r *EncargoGORM) Actualizar(a *models.Encargo) error {
	// TODO: implementar.
	return errors.New("TODO: implementar Actualizar")
}
