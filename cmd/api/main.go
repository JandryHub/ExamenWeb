// ARCHIVO BLOQUEADO — NO MODIFICAR
package main

import (
	"log"
	"net/http"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/joancema/examen-floristeria/internal/handlers"
	"github.com/joancema/examen-floristeria/internal/models"
	"github.com/joancema/examen-floristeria/internal/services"
	"github.com/joancema/examen-floristeria/internal/storage"
)

func main() {
	db, err := gorm.Open(sqlite.Open("floristeria.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("no se pudo abrir la base de datos: %v", err)
	}

	if err := db.AutoMigrate(
		&models.Arreglo{},
		&models.Cliente{},
		&models.Encargo{},
	); err != nil {
		log.Fatalf("error en la migración: %v", err)
	}

	sembrarArreglos(db)

	// Repositories (GORM)
	arregloRepo := storage.NuevoArregloGORM(db)
	clienteRepo := storage.NuevoClienteGORM(db)
	encargoRepo := storage.NuevoEncargoGORM(db)

	// Services
	arregloSvc := services.NuevoArregloService(arregloRepo)
	clienteSvc := services.NuevoClienteService(clienteRepo)
	encargoSvc := services.NuevoEncargoService(encargoRepo, arregloRepo, clienteRepo)

	// Handlers + Router
	router := handlers.NuevoRouter(
		handlers.NuevoArregloHandler(arregloSvc),
		handlers.NuevoClienteHandler(clienteSvc),
		handlers.NuevoEncargoHandler(encargoSvc),
	)

	log.Println("API de la floristería escuchando en http://localhost:8082")
	if err := http.ListenAndServe(":8082", router); err != nil {
		log.Fatal(err)
	}
}

// sembrarArreglos carga el catálogo inicial solo si la tabla está vacía.
// Los clientes y encargos se crean vía API.
func sembrarArreglos(db *gorm.DB) {
	var total int64
	db.Model(&models.Arreglo{}).Count(&total)
	if total > 0 {
		return
	}
	iniciales := []models.Arreglo{
		{Nombre: "Ramo de rosas", PrecioUnitario: 8.50, Stock: 10, Activo: true},
		{Nombre: "Ramo de girasoles", PrecioUnitario: 6.00, Stock: 4, Activo: true},
		{Nombre: "Centro de mesa", PrecioUnitario: 5.00, Stock: 2, Activo: true},
		{Nombre: "Arco floral", PrecioUnitario: 15.00, Stock: 3, Activo: false},
	}
	for i := range iniciales {
		db.Create(&iniciales[i])
	}
}
