package sqlwrapper

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 15432
	user     = "postgres"
	password = "mypassword"
	dbname   = "postgres"
)

// ConnectToPostgres é uma função que irá se conectar com o postgres.
func ConnectToPostgres() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")
	return db
}

func InsertIntoTable(db *sql.DB, tableName string, values string) {
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	query := fmt.Sprintf(`INSERT INTO compras (cpf, private, incompleto, dataUltimaCompra, ticketMedio, ticketUltimaCompra, lojaMaisFrequente, lojaUltimaCompra) VALUES (%s);`, values)
	fmt.Println(query)

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully inserted!")
}
