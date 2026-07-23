// ARCHIVO BLOQUEADO — NO MODIFICAR
package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/joancema/examen-floristeria/internal/models"
	"github.com/joancema/examen-floristeria/internal/services"
)

// ArregloHandler expone la Entidad A por HTTP.
// Está completo: observe cómo decodifica el body, llama al service y
// MAPEA los errores de dominio a status codes. Ese mapeo es exactamente
// lo que usted debe replicar en sus propios handlers.
type ArregloHandler struct {
	servicio *services.ArregloService
}

func NuevoArregloHandler(s *services.ArregloService) *ArregloHandler {
	return &ArregloHandler{servicio: s}
}

func (h *ArregloHandler) Listar(w http.ResponseWriter, r *http.Request) {
	lista, err := h.servicio.Listar()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, lista)
}

func (h *ArregloHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var arreglo models.Arreglo
	if err := json.NewDecoder(r.Body).Decode(&arreglo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido")
		return
	}
	if err := h.servicio.Crear(&arreglo); err != nil {
		switch {
		case errors.Is(err, services.ErrDatosInvalidos):
			RespondError(w, http.StatusUnprocessableEntity, err.Error())
		default:
			RespondError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	RespondJSON(w, http.StatusCreated, arreglo)
}
