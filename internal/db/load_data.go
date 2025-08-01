package db

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"dark_matter_stock/internal/api"
)

func Migrate(ctx context.Context) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS stocks (
		ticker VARCHAR(255) PRIMARY KEY,
		company VARCHAR(255),
		brokerage VARCHAR(255),
		action VARCHAR(255),
		target_from DECIMAL,
		target_to DECIMAL,
		rating_from VARCHAR(255),
		rating_to VARCHAR(255),
		time TIMESTAMP
	);`

	_, err := Conn.Exec(ctx, createTableSQL)
	if err != nil {
		return err
	}
	log.Println("Tabla 'stocks' creada o ya existe.")
	return nil
}

func cleanAndParseFloatPtr(s string) *float64 {

	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "$", "")
	s = strings.ReplaceAll(s, ",", "")

	if s == "" {
		return nil
	}

	val, err := strconv.ParseFloat(s, 64)
	if err != nil {

		log.Printf("Error de parsing para el valor '%s': %v", s, err)
		return nil
	}

	return &val
}

func InsertStocks(ctx context.Context, items []api.StockItem) error {
	tx, err := Conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, item := range items {

		targetFrom := cleanAndParseFloatPtr(item.TargetFrom)
		targetTo := cleanAndParseFloatPtr(item.TargetTo)
		ratingFrom := item.RatingFrom
		ratingTo := item.RatingTo
		action := item.Action

		t, err := time.Parse(time.RFC3339Nano, item.Time)
		if err != nil {
			log.Printf("Error al parsear la fecha para %s: %v", item.Ticker, err)
			continue
		}

		_, err = tx.Exec(ctx, "UPSERT INTO stocks (ticker, company, brokerage, action, target_from, target_to, rating_from, rating_to, time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
			item.Ticker, item.Company, item.Brokerage, action, targetFrom, targetTo, ratingFrom, ratingTo, t)
		if err != nil {
			log.Printf("Error al insertar %s: %v", item.Ticker, err)
			return err
		}
	}

	return tx.Commit(ctx)
}
