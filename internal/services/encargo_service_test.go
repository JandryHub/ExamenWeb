package services

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/joancema/examen-floristeria/internal/models"
	"github.com/joancema/examen-floristeria/internal/storage"
)

// Requisito CP3: 2 Tests propios.

func TestPropio_1_Descuento(t *testing.T) {
	arreglos := storage.NuevoArregloMemoria()
	clientes := storage.NuevoClienteMemoria()
	encargos := storage.NuevoEncargoMemoria()

	svc := NuevoEncargoService(encargos, arreglos, clientes)

	arreglo := models.Arreglo{Model: gorm.Model{ID: 1}, Nombre: "Ramo VIP", PrecioUnitario: 10.0, Stock: 20, Activo: true}
	arreglos.Crear(&arreglo)

	cliente := models.Cliente{Model: gorm.Model{ID: 1}, Nombre: "Test", Cedula: "123", Telefono: "123"}
	clientes.Crear(&cliente)

	// Comprar 5 unidades para activar el descuento del 10%
	encargo := models.Encargo{
		ArregloID: 1,
		ClienteID: 1,
		Cantidad:  5,
	}
	err := svc.Crear(&encargo)

	require.NoError(t, err)
	// 5 * 10 = 50. Menos 10% = 45.0
	require.Equal(t, 45.0, encargo.Total, "Debería aplicar 10% de descuento exacto")
	require.Equal(t, models.EstadoPendiente, encargo.Estado)
}

func TestPropio_2_Cancelacion_Invalida(t *testing.T) {
	arreglos := storage.NuevoArregloMemoria()
	clientes := storage.NuevoClienteMemoria()
	encargos := storage.NuevoEncargoMemoria()

	svc := NuevoEncargoService(encargos, arreglos, clientes)

	// Insertamos un encargo directamente al repositorio simulando que ya fue ENTREGADO
	encargo := models.Encargo{
		ArregloID: 1,
		ClienteID: 1,
		Cantidad:  1,
		Estado:    models.EstadoEntregado,
	}
	encargos.Crear(&encargo)

	// Intentamos cancelar un encargo que ya fue entregado
	err := svc.Cancelar(encargo.ID)

	require.Error(t, err)
	require.ErrorIs(t, err, ErrEstadoInvalido, "No se debe poder cancelar un encargo que no esté en estado PENDIENTE")
}
