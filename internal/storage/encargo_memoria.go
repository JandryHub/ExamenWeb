// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import (
	"sync"

	"github.com/joancema/examen-floristeria/internal/models"
)

// EncargoMemoria implementa EncargoRepository en memoria.
// Se usa en los tests de reglas de negocio como fake del repositorio real.
type EncargoMemoria struct {
	mu     sync.Mutex
	datos  map[uint]models.Encargo
	nextID uint
}

func NuevoEncargoMemoria() *EncargoMemoria {
	return &EncargoMemoria{datos: make(map[uint]models.Encargo), nextID: 1}
}

func (r *EncargoMemoria) Crear(a *models.Encargo) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	a.ID = r.nextID
	r.nextID++
	r.datos[a.ID] = *a
	return nil
}

func (r *EncargoMemoria) ObtenerPorID(id uint) (models.Encargo, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	a, ok := r.datos[id]
	return a, ok
}

func (r *EncargoMemoria) Listar() ([]models.Encargo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	lista := make([]models.Encargo, 0, len(r.datos))
	for _, a := range r.datos {
		lista = append(lista, a)
	}
	return lista, nil
}

func (r *EncargoMemoria) Actualizar(a *models.Encargo) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.datos[a.ID]; !ok {
		return ErrRegistroNoExiste
	}
	r.datos[a.ID] = *a
	return nil
}
