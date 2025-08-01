package main

import (
	"context"
	"log"

	"github.com/DarkMatterStock/internal/api"
	"github.com/DarkMatterStock/internal/db"
)

func main() {
	ctx := context.Background()
	log.Println("Iniciando la aplicación DarkMatterStock...")

	// 1. Conectar a la base de datos
	err := db.InitDB(ctx)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Conn.Close(ctx)

	// 2. Crear la tabla 'stocks' si no existe
	err = db.Migrate(ctx)
	if err != nil {
		log.Fatalf("Error al migrar la base de datos: %v", err)
	}

	// 3. Descargar todos los stocks de la API
	log.Println("Comenzando a descargar datos de la API...")
	stocks, err := api.DownloadAllStocks()
	if err != nil {
		log.Fatalf("Error al descargar datos de la API: %v", err)
	}
	log.Printf("Descargados %d items de stocks de la API.", len(stocks))

	// 4. Insertar los stocks en la base de datos
	log.Println("Insertando datos en la base de datos...")
	err = db.InsertStocks(ctx, stocks)
	if err != nil {
		log.Fatalf("Error al insertar datos en la base de datos: %v", err)
	}
	log.Println("Proceso de carga de datos finalizado con éxito.")
}
