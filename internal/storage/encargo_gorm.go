package storage

import (
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
	return r.db.Create(a).Error
}

func (r *EncargoGORM) ObtenerPorID(id uint) (models.Encargo, bool) {
	var a models.Encargo
	result := r.db.Preload("Arreglo").Preload("Cliente").First(&a, id)
	return a, result.Error == nil
}

func (r *EncargoGORM) Listar() ([]models.Encargo, error) {
	var encargos []models.Encargo
	err := r.db.Preload("Arreglo").Preload("Cliente").Find(&encargos).Error
	return encargos, err
}

func (r *EncargoGORM) Actualizar(a *models.Encargo) error {
	return r.db.Save(a).Error
}
