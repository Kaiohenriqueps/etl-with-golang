package csvwrapper

import (
	"bufio"
	"etl-with-golang/src/sqlwrapper"
	"fmt"
	"log"
	"os"
	"strings"
)

// MyStruct é a estrutura do item de cada linha.
type MyStruct struct {
	cpf                string
	private            string
	incompleto         string
	dataUltimaCompra   string
	ticketMedio        string
	ticketUltimaCompra string
	lojaMaisFrequente  string
	lojaUltimaCompra   string
}

// OpenFile é uma função que lê um arquivo e retorna o arquivo em si.
func OpenFile(path string) []MyStruct {
	csvfile, err := os.Open(path)
	if err != nil {
		log.Println("Não conseguiu abrir o arquivo")
		log.Fatal(err)
	}
	log.Println("Abriu o arquivo com sucesso!")

	fileScanner := bufio.NewScanner(csvfile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	csvfile.Close()

	return CreateStruct(fileTextLines)
}

// CreateStruct é uma função que cria um array de MyStruct.
func CreateStruct(fileTextLines []string) []MyStruct {
	var objs []MyStruct

	for _, line := range fileTextLines {
		fields := strings.Fields(line)
		if fields[0] == "CPF" {
			continue
		}
		item := MyStruct{
			cpf:                fields[0],
			private:            fields[1],
			incompleto:         fields[2],
			dataUltimaCompra:   fields[3],
			ticketMedio:        fields[4],
			ticketUltimaCompra: fields[5],
			lojaMaisFrequente:  fields[6],
			lojaUltimaCompra:   fields[7],
		}

		objs = append(objs, item)
	}
	return objs
}

// InsertStructs é uma função que insere as MyStructs na tabela do postgres.
func InsertStructs(path string) {
	objs := OpenFile(path)

	conn := sqlwrapper.ConnectToPostgres()

	for _, item := range objs {
		values := fmt.Sprintf("'%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s'",
			item.cpf, item.private, item.incompleto, item.dataUltimaCompra, item.ticketMedio,
			item.ticketUltimaCompra, item.lojaMaisFrequente, item.lojaUltimaCompra)
		sqlwrapper.InsertIntoTable(conn, "compras", values)
	}
}
