package db

import (
	"context"

	"github.com/ssanchezgo/DarkMatterStock/internal/api"
)

// GetAllStocks consulta y devuelve todos los registros de la tabla 'stocks'.
func GetAllStocks(ctx context.Context) ([]api.StockItem, error) {
	rows, err := Conn.Query(ctx, "SELECT * FROM stocks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stocks []api.StockItem
	for rows.Next() {
		var s api.StockItem
		err := rows.Scan(
			&s.Ticker,
			&s.Company,
			&s.Brokerage,
			&s.Action,
			&s.TargetFrom,
			&s.TargetTo,
			&s.RatingFrom,
			&s.RatingTo,
			&s.Time,
		)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, s)
	}

	return stocks, rows.Err()
}

// GetStockByTicker consulta y devuelve un registro de la tabla 'stocks' por su ticker.
func GetStockByTicker(ctx context.Context, ticker string) (api.StockItem, error) {
	var s api.StockItem
	err := Conn.QueryRow(ctx, "SELECT * FROM stocks WHERE ticker = $1", ticker).Scan(
		&s.Ticker,
		&s.Company,
		&s.Brokerage,
		&s.Action,
		&s.TargetFrom,
		&s.TargetTo,
		&s.RatingFrom,
		&s.RatingTo,
		&s.Time,
	)
	if err != nil {
		return s, err
	}
	return s, nil
}
