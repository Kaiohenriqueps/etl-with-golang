package sqlwrapper

import (
	"database/sql"
	"etl-with-golang/src/utils"
	"fmt"
	"log"
	"os"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

// Colunas da tabela
var columns = []string{
	"cpf", "private", "incompleto", "dataUltimaCompra",
	"ticketMedio", "ticketUltimaCompra", "lojaMaisFrequente",
	"lojaUltimaCompra", "flagCPF", "flagCNPJFrequente", "flagCNPJUltima"}

// ConnectToPostgres é uma função que irá se conectar com o postgres.
func ConnectToPostgres() *sql.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASS"), os.Getenv("DBNAME"))

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
	txn, err := conn.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := txn.Prepare(pq.CopyIn("compras", "cpf", "private", "incompleto", "dataultimacompra", "ticketmedio", "ticketultimacompra",
		"lojamaisfrequente", "lojaultimacompra", "flagcpf", "flagcnpjfrequente", "flagcnpjultima"))

	for _, item := range objs {
		_, err = stmt.Exec(string(item.Cpf), string(item.Private), string(item.Incompleto), string(item.DataUltimaCompra),
			string(item.TicketMedio), string(item.TicketUltimaCompra), string(item.LojaMaisFrequente), string(item.LojaUltimaCompra),
			string(item.FlagCPF), string(item.FlagCNPJFrequente), string(item.FlagCNPJUltima))

		if err != nil {
			log.Fatal(err)
		}
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
	}

}
