package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Conectar() {
	connStr := "host=localhost port=5432 user=postgres password=neagu1234 dbname=leads_db sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao abrir conexão com banco:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Não foi possível conectar ao banco:", err)
	}

	fmt.Println("✅ Conectado ao PostgreSQL com sucesso!")

}
