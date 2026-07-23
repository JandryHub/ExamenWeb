// ARCHIVO BLOQUEADO — NO MODIFICAR
//
// Las 5 reglas de negocio se verifican aquí usando los repositorios EN MEMORIA
// (ya implementados en el repo base) como fakes. Así, estos tests miden solo
// la lógica de su EncargoService, sin depender de su implementación GORM.
package acceptance

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/joancema/examen-floristeria/internal/models"
	"github.com/joancema/examen-floristeria/internal/services"
	"github.com/joancema/examen-floristeria/internal/storage"
)

type entornoReglas struct {
	svc          *services.EncargoService
	arreglos *storage.ArregloMemoria
	clientes     *storage.ClienteMemoria
	encargos   *storage.EncargoMemoria
	principal      models.Arreglo
	ana          models.Cliente
}

func nuevoEntornoReglas(t *testing.T) entornoReglas {
	t.Helper()
	hm := storage.NuevoArregloMemoria()
	cm := storage.NuevoClienteMemoria()
	am := storage.NuevoEncargoMemoria()

	principal := models.Arreglo{Nombre: "Ramo de rosas", PrecioUnitario: 8.5, Stock: 10, Activo: true}
	require.NoError(t, hm.Crear(&principal))
	ana := models.Cliente{Nombre: "Ana Zambrano", Cedula: "1310000001", Telefono: "0990000001"}
	require.NoError(t, cm.Crear(&ana))

	return entornoReglas{
		svc:          services.NuevoEncargoService(am, hm, cm),
		arreglos: hm,
		clientes:     cm,
		encargos:   am,
		principal:      principal,
		ana:          ana,
	}
}

// R1: no se crea un encargo si el arreglo no existe o está inactivo,
// o si el cliente no existe.
func TestCP2_R1_ReferenciasValidas(t *testing.T) {
	e := nuevoEntornoReglas(t)

	a := models.Encargo{ArregloID: 99999, ClienteID: e.ana.ID, Cantidad: 1}
	require.ErrorIs(t, e.svc.Crear(&a), services.ErrReferenciaInvalida,
		"crear con un arreglo inexistente debe devolver ErrReferenciaInvalida")

	extra := models.Arreglo{Nombre: "Arco floral", PrecioUnitario: 15, Stock: 3, Activo: false}
	require.NoError(t, e.arreglos.Crear(&extra))
	a = models.Encargo{ArregloID: extra.ID, ClienteID: e.ana.ID, Cantidad: 1}
	require.ErrorIs(t, e.svc.Crear(&a), services.ErrReferenciaInvalida,
		"crear con un arreglo INACTIVO debe devolver ErrReferenciaInvalida")

	a = models.Encargo{ArregloID: e.principal.ID, ClienteID: 99999, Cantidad: 1}
	require.ErrorIs(t, e.svc.Crear(&a), services.ErrReferenciaInvalida,
		"crear con un cliente inexistente debe devolver ErrReferenciaInvalida")
}

// R2: la cantidad no puede superar el stock disponible.
func TestCP2_R2_StockInsuficiente(t *testing.T) {
	e := nuevoEntornoReglas(t)

	a := models.Encargo{ArregloID: e.principal.ID, ClienteID: e.ana.ID, Cantidad: 11}
	require.ErrorIs(t, e.svc.Crear(&a), services.ErrStockInsuficiente,
		"pedir 11 unidades con stock 10 debe devolver ErrStockInsuficiente")
}

// R3: Total = Cantidad x PrecioUnitario, con 10% de descuento desde 5 unidades.
func TestCP2_R3_CalculoTotal(t *testing.T) {
	e := nuevoEntornoReglas(t)

	sinDescuento := models.Encargo{ArregloID: e.principal.ID, ClienteID: e.ana.ID, Cantidad: 3}
	require.NoError(t, e.svc.Crear(&sinDescuento),
		"crear un encargo válido no debe devolver error")
	require.InDelta(t, 25.50, sinDescuento.Total, 0.001,
		"3 x 8.50 = 25.50 (sin descuento)")
	require.Equal(t, models.EstadoPendiente, sinDescuento.Estado,
		"un encargo recién creado debe quedar en estado PENDIENTE")

	conDescuento := models.Encargo{ArregloID: e.principal.ID, ClienteID: e.ana.ID, Cantidad: 5}
	require.NoError(t, e.svc.Crear(&conDescuento))
	require.InDelta(t, 38.25, conDescuento.Total, 0.001,
		"5 x 8.50 = 42.50, con 10% de descuento = 38.25")
}

// R4: solo se puede cancelar un encargo en estado PENDIENTE.
func TestCP2_R4_CancelarSoloPendiente(t *testing.T) {
	e := nuevoEntornoReglas(t)

	entregado := models.Encargo{
		ArregloID: e.principal.ID,
		ClienteID:     e.ana.ID,
		Cantidad:      1,
		Estado:        models.EstadoEntregado,
		Total:         8.5,
	}
	require.NoError(t, e.encargos.Crear(&entregado))
	require.ErrorIs(t, e.svc.Cancelar(entregado.ID), services.ErrEstadoInvalido,
		"cancelar un encargo ENTREGADO debe devolver ErrEstadoInvalido")

	require.ErrorIs(t, e.svc.Cancelar(99999), services.ErrNoEncontrado,
		"cancelar un encargo inexistente debe devolver ErrNoEncontrado")
}

// R5: al crear se descuenta el stock; al cancelar, se repone.
func TestCP2_R5_ReposicionStock(t *testing.T) {
	e := nuevoEntornoReglas(t)

	a := models.Encargo{ArregloID: e.principal.ID, ClienteID: e.ana.ID, Cantidad: 3}
	require.NoError(t, e.svc.Crear(&a))

	h, ok := e.arreglos.ObtenerPorID(e.principal.ID)
	require.True(t, ok)
	require.Equal(t, uint(7), h.Stock,
		"al crear un encargo de 3 unidades, el stock debe bajar de 10 a 7")

	require.NoError(t, e.svc.Cancelar(a.ID), "cancelar un encargo PENDIENTE debe funcionar")

	cancelado, ok := e.encargos.ObtenerPorID(a.ID)
	require.True(t, ok)
	require.Equal(t, models.EstadoCancelado, cancelado.Estado,
		"tras cancelar, el encargo debe quedar en estado CANCELADO")

	h, ok = e.arreglos.ObtenerPorID(e.principal.ID)
	require.True(t, ok)
	require.Equal(t, uint(10), h.Stock,
		"al cancelar, las 3 unidades deben reponerse al stock (7 -> 10)")
}
