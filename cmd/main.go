package main

import (
	"context"
	"log"

	"dark_matter_stock/internal/db"
	"dark_matter_stock/internal/server"
)

func main() {
	ctx := context.Background()
	log.Println("Iniciando la aplicación DarkMatterStock...")

	err := db.InitDB(ctx)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Conn.Close(ctx)
	log.Println("Conexión a CockroachDB exitosa!")

	err = db.Migrate(ctx)
	if err != nil {
		log.Fatalf("Error al migrar la base de datos: %v", err)
	}
	log.Println("Tabla 'stocks' creada o ya existe.")

	apiServer := server.NewServer()
	if err := apiServer.Run("8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
