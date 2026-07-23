// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import (
	"sync"

	"github.com/joancema/examen-floristeria/internal/models"
)

// ArregloMemoria implementa ArregloRepository en memoria.
// Se usa en los tests de reglas de negocio como fake del repositorio real.
type ArregloMemoria struct {
	mu     sync.Mutex
	datos  map[uint]models.Arreglo
	nextID uint
}

func NuevoArregloMemoria() *ArregloMemoria {
	return &ArregloMemoria{datos: make(map[uint]models.Arreglo), nextID: 1}
}

func (r *ArregloMemoria) Crear(h *models.Arreglo) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	h.ID = r.nextID
	r.nextID++
	r.datos[h.ID] = *h
	return nil
}

func (r *ArregloMemoria) ObtenerPorID(id uint) (models.Arreglo, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	h, ok := r.datos[id]
	return h, ok
}

func (r *ArregloMemoria) Listar() ([]models.Arreglo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	lista := make([]models.Arreglo, 0, len(r.datos))
	for _, h := range r.datos {
		lista = append(lista, h)
	}
	return lista, nil
}

func (r *ArregloMemoria) Actualizar(h *models.Arreglo) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.datos[h.ID]; !ok {
		return ErrRegistroNoExiste
	}
	r.datos[h.ID] = *h
	return nil
}
