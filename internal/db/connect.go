package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func InitDB(ctx context.Context) error {

	dsn := "postgresql://ssanchezgo:LDysEX3GNazGiy-jsexKKg@dark-matter-stock-02-13155.j77.aws-us-east-1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
	var err error
	Conn, err = pgx.Connect(ctx, dsn)
	if err != nil {
		return err
	}
	err = Conn.Ping(ctx)
	if err != nil {
		Conn.Close(ctx)
		return err
	}

	log.Println("Conexión a CockroachDB exitosa!")
	return nil
}
