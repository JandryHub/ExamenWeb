package services

import (
	"github.com/joancema/examen-floristeria/internal/models"
	"github.com/joancema/examen-floristeria/internal/storage"
)

// TAREA (CP2): Implemente EncargoService con las 5 reglas de negocio.
//
// Las reglas están A LA VISTA en las pantallas (carpeta pantallas/) y los
// tests de acceptance/reglas_test.go las verifican una por una. Devuelva los
// errores de dominio de errores.go: los tests los comprueban con errors.Is.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos.
//   - Observe que el service recibe TRES repositories: necesita consultar
//     Arreglo y Cliente para validar, y actualizar Arreglo al cancelar.
type EncargoService struct {
	encargos   storage.EncargoRepository
	arreglos storage.ArregloRepository
	clientes     storage.ClienteRepository
}

func NuevoEncargoService(
	encargos storage.EncargoRepository,
	arreglos storage.ArregloRepository,
	clientes storage.ClienteRepository,
) *EncargoService {
	return &EncargoService{
		encargos:   encargos,
		arreglos: arreglos,
		clientes:     clientes,
	}
}

// Crear registra un nuevo encargo aplicando R1, R2 y R3.
// TODO (R1): el arreglo debe existir y estar activo; el cliente debe existir.
// TODO (R2): la cantidad no puede superar el stock disponible del arreglo.
// TODO (R3): calcule el total (observe en las pantallas cuándo aplica descuento).
// TODO: al crear, el stock del arreglo se descuenta (mire la pantalla 01
// antes y después de crear un encargo; R5 es la operación inversa).
func (s *EncargoService) Crear(a *models.Encargo) error {
	arreglo, ok := s.arreglos.ObtenerPorID(a.ArregloID)
	if !ok || !arreglo.Activo {
		return ErrReferenciaInvalida
	}

	_, ok = s.clientes.ObtenerPorID(a.ClienteID)
	if !ok {
		return ErrReferenciaInvalida
	}

	if a.Cantidad > arreglo.Stock {
		return ErrStockInsuficiente
	}

	total := float64(a.Cantidad) * arreglo.PrecioUnitario
	if a.Cantidad >= 5 {
		total = total * 0.9 // 10% de descuento
	}
	a.Total = total
	a.Estado = models.EstadoPendiente

	arreglo.Stock -= a.Cantidad
	if err := s.arreglos.Actualizar(&arreglo); err != nil {
		return err
	}

	return s.encargos.Crear(a)
}

func (s *EncargoService) ObtenerPorID(id uint) (models.Encargo, error) {
	encargo, ok := s.encargos.ObtenerPorID(id)
	if !ok {
		return models.Encargo{}, ErrNoEncontrado
	}
	return encargo, nil
}

func (s *EncargoService) Listar() ([]models.Encargo, error) {
	return s.encargos.Listar()
}

// Cancelar cancela un encargo aplicando R4 y R5.
// TODO (R4): solo se puede cancelar un encargo en estado PENDIENTE.
// TODO (R5): al cancelar, la cantidad se repone al stock del arreglo.
func (s *EncargoService) Cancelar(id uint) error {
	encargo, ok := s.encargos.ObtenerPorID(id)
	if !ok {
		return ErrNoEncontrado
	}
	if encargo.Estado != models.EstadoPendiente {
		return ErrEstadoInvalido
	}

	arreglo, ok := s.arreglos.ObtenerPorID(encargo.ArregloID)
	if ok {
		arreglo.Stock += encargo.Cantidad
		s.arreglos.Actualizar(&arreglo)
	}

	encargo.Estado = models.EstadoCancelado
	return s.encargos.Actualizar(&encargo)
}
