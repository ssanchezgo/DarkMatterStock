package db

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/DarkMatterStock/internal/api"
)

// Migrate crea la tabla 'stocks' si no existe.
func Migrate(ctx context.Context) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS stocks (
		ticker VARCHAR(255) PRIMARY KEY,
		company VARCHAR(255),
		brokerage VARCHAR(255),
		action VARCHAR(255),
		target_from DECIMAL,
		target_to DECIMAL,
		rating_from DECIMAL,
		rating_to DECIMAL,
		time TIMESTAMP
	);`

	_, err := Conn.Exec(ctx, createTableSQL)
	if err != nil {
		return err
	}
	log.Println("Tabla 'stocks' creada o ya existe.")
	return nil
}

// InsertStocks inserta una lista de acciones en la base de datos.
func InsertStocks(ctx context.Context, items []api.StockItem) error {
	tx, err := Conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Prepara la sentencia para la inserción
	insertStmt, err := tx.Prepare(ctx, "insert_stocks", "UPSERT INTO stocks (ticker, company, brokerage, action, target_from, target_to, rating_from, rating_to, time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)")
	if err != nil {
		return err
	}

	for _, item := range items {
		// Conversión de strings a float64, manejando errores de parsing.
		targetFrom, err := strconv.ParseFloat(item.TargetFrom, 64)
		if err != nil {
			targetFrom = 0 // o manejar el error de otra forma, por ejemplo, ignorando el campo
		}

		targetTo, err := strconv.ParseFloat(item.TargetTo, 64)
		if err != nil {
			targetTo = 0
		}

		ratingFrom, err := strconv.ParseFloat(item.RatingFrom, 64)
		if err != nil {
			ratingFrom = 0
		}

		ratingTo, err := strconv.ParseFloat(item.RatingTo, 64)
		if err != nil {
			ratingTo = 0
		}

		t, err := time.Parse("2006-01-02 15:04:05", item.Time)
		if err != nil {
			log.Printf("Error al parsear la fecha para %s: %v", item.Ticker, err)
			continue
		}

		_, err = tx.Exec(ctx, "insert_stocks", item.Ticker, item.Company, item.Brokerage, item.Action, targetFrom, targetTo, ratingFrom, ratingTo, t)
		if err != nil {
			log.Printf("Error al insertar %s: %v", item.Ticker, err)
			return err
		}
	}

	return tx.Commit(ctx)
}
