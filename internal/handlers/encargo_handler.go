package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/joancema/examen-floristeria/internal/models"
	"github.com/joancema/examen-floristeria/internal/services"
)

// TAREA (CP3): Implemente EncargoHandler.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos.
//   - Mapeo de errores de dominio a status codes (los tests lo verifican):
//       ErrDatosInvalidos     -> 422 Unprocessable Entity
//       ErrReferenciaInvalida -> 422 Unprocessable Entity
//       ErrStockInsuficiente  -> 409 Conflict
//       ErrEstadoInvalido     -> 409 Conflict
//       ErrNoEncontrado       -> 404 Not Found
//       cualquier otro error  -> 500 Internal Server Error
type EncargoHandler struct {
	servicio *services.EncargoService
}

func NuevoEncargoHandler(s *services.EncargoService) *EncargoHandler {
	return &EncargoHandler{servicio: s}
}

func (h *EncargoHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var encargo models.Encargo
	if err := json.NewDecoder(r.Body).Decode(&encargo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	if err := h.servicio.Crear(&encargo); err != nil {
		switch {
		case errors.Is(err, services.ErrDatosInvalidos) || errors.Is(err, services.ErrReferenciaInvalida):
			RespondError(w, http.StatusUnprocessableEntity, err.Error())
		case errors.Is(err, services.ErrStockInsuficiente):
			RespondError(w, http.StatusConflict, err.Error())
		default:
			RespondError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	RespondJSON(w, http.StatusCreated, encargo)
}

func (h *EncargoHandler) Listar(w http.ResponseWriter, r *http.Request) {
	lista, err := h.servicio.Listar()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, lista)
}

func (h *EncargoHandler) ObtenerPorID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	encargo, err := h.servicio.ObtenerPorID(uint(id))
	if err != nil {
		switch {
		case errors.Is(err, services.ErrNoEncontrado):
			RespondError(w, http.StatusNotFound, err.Error())
		default:
			RespondError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	RespondJSON(w, http.StatusOK, encargo)
}

func (h *EncargoHandler) Cancelar(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	if err := h.servicio.Cancelar(uint(id)); err != nil {
		switch {
		case errors.Is(err, services.ErrNoEncontrado):
			RespondError(w, http.StatusNotFound, err.Error())
		case errors.Is(err, services.ErrEstadoInvalido):
			RespondError(w, http.StatusConflict, err.Error())
		default:
			RespondError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	RespondJSON(w, http.StatusOK, map[string]string{"mensaje": "cancelado con éxito"})
}
