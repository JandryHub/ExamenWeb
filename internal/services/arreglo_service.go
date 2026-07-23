// ARCHIVO BLOQUEADO — NO MODIFICAR
package services

import (
	"github.com/joancema/examen-floristeria/internal/models"
	"github.com/joancema/examen-floristeria/internal/storage"
)

// ArregloService contiene la lógica de negocio de la Entidad A.
// Está completo: úselo como ejemplo de cómo un service valida datos,
// devuelve errores de dominio y delega la persistencia al repository.
type ArregloService struct {
	repo storage.ArregloRepository
}

func NuevoArregloService(repo storage.ArregloRepository) *ArregloService {
	return &ArregloService{repo: repo}
}

func (s *ArregloService) Crear(h *models.Arreglo) error {
	if h.Nombre == "" || h.PrecioUnitario <= 0 {
		return ErrDatosInvalidos
	}
	return s.repo.Crear(h)
}

func (s *ArregloService) ObtenerPorID(id uint) (models.Arreglo, error) {
	h, ok := s.repo.ObtenerPorID(id)
	if !ok {
		return models.Arreglo{}, ErrNoEncontrado
	}
	return h, nil
}

func (s *ArregloService) Listar() ([]models.Arreglo, error) {
	return s.repo.Listar()
}
