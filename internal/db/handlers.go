// package db

// import (
// 	"encoding/json"
// 	"net/http"

// 	"yourmodulepath/model"

// 	"github.com/jackc/pgx/v4"
// )

// // HandleGetStocks handles the HTTP request to get stock items from the database.
// func HandleGetStocks(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	dsn := "postgres://user:password@localhost:5432/dbname" // TODO: Replace with your actual DSN or load from config
// 	conn, err := pgx.Connect(ctx, dsn)
// 	if err != nil {
// 		http.Error(w, "DB Connection Error", http.StatusInternalServerError)
// 		return
// 	}
// 	defer conn.Close(ctx)

// 	rows, err := conn.Query(ctx, "SELECT ticker, company, brokerage, action, target_from, target_to, rating_from, rating_to, time FROM stock")
// 	if err != nil {
// 		http.Error(w, "Query Error", http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	var stocks []model.StockItem
// 	for rows.Next() {
// 		var s model.StockItem
// 		if err := rows.Scan(&s.Ticker, &s.Company, &s.Brokerage, &s.Action, &s.TargetFrom, &s.TargetTo, &s.RatingFrom, &s.RatingTo, &s.Time); err != nil {
// 			http.Error(w, "Scan Error", http.StatusInternalServerError)
// 			return
// 		}
// 		stocks = append(stocks, s)
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		if err := json.NewEncoder(w).Encode(stocks); err != nil {
// 			http.Error(w, "JSON Encoding Error", http.StatusInternalServerError)
// 			return
// 		}
// 		return
// 	}
// }
