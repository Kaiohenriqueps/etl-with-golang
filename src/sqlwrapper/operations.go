package sqlwrapper

import (
	"database/sql"
	"etl-with-golang/src/utils"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mypassword"
	dbname   = "postgres"
)

// ConnectToPostgres é uma função que irá se conectar com o postgres.
func ConnectToPostgres() *sql.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// InsertStructs é uma função que insere as MyStructs na tabela do postgres.
// @param objs: array de MyStructs.
func InsertStructs(objs []utils.MyStruct) {
	conn := ConnectToPostgres()

	for _, item := range objs {
		values := fmt.Sprintf("'%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s'",
			item.Cpf, item.Private, item.Incompleto, item.DataUltimaCompra, item.TicketMedio,
			item.TicketUltimaCompra, item.LojaMaisFrequente, item.LojaUltimaCompra)
		InsertIntoTable(conn, "compras", values)
	}

}

// InsertIntoTable é uma função que insere as MyStructs na tabela do postgres.
// @param db: conexão do banco de dados.
// @param tableName: nome da tabela que será populada.
// @param values: valores que serão inseridos na tabela.
func InsertIntoTable(db *sql.DB, tableName string, values string) {
	// db.SetMaxIdleConns(150)
	// db.SetMaxOpenConns(150)

	query := fmt.Sprintf(`INSERT INTO %s (cpf, private, incompleto, dataUltimaCompra, `+
		`ticketMedio, ticketUltimaCompra, lojaMaisFrequente, lojaUltimaCompra) VALUES (%s);`,
		tableName, values)

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
