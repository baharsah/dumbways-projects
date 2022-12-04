package conn

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func DatabaseConnect() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	databaseUrl := "postgres://postgres:baharsahtest@localhost:5432/postgres"

	var err error
	Conn, err = pgx.Connect(context.Background(), databaseUrl)

	if err != nil {
		log.Println("Koneksi database gagal", err)
	}

	fmt.Println("Koneksi ke database berhasil!!")
}
