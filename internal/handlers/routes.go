// ARCHIVO BLOQUEADO — NO MODIFICAR
package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NuevoRouter registra todas las rutas de la API. Este archivo es el
// contrato HTTP del examen: los tests httptest de acceptance/ atacan
// exactamente estas rutas.
func NuevoRouter(
	arreglos *ArregloHandler,
	clientes *ClienteHandler,
	encargos *EncargoHandler,
) http.Handler {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/arreglos", func(r chi.Router) {
			r.Get("/", arreglos.Listar)
			r.Post("/", arreglos.Crear)
		})

		r.Route("/clientes", func(r chi.Router) {
			r.Get("/", clientes.Listar)
			r.Post("/", clientes.Crear)
			r.Get("/{id}", clientes.ObtenerPorID)
		})

		r.Route("/encargos", func(r chi.Router) {
			r.Get("/", encargos.Listar)
			r.Post("/", encargos.Crear)
			r.Get("/{id}", encargos.ObtenerPorID)
			r.Post("/{id}/cancelar", encargos.Cancelar)
		})
	})

	return r
}
