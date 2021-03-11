package csvwrapper

import (
	"bufio"
	"etl-with-golang/src/utils"
	"log"
	"os"
	"strings"
)

// OpenFileAndCreateStruct é uma função que lê um arquivo e retorna o arquivo em si.
// @param path: caminho do arquivo que será processado.
func OpenFileAndCreateStruct(path string) []utils.MyStruct {
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

// VerificaDocs é uma função que verifica se os documentos são válidos.
// @param item: struct do registro.
func VerificaDocs(item utils.MyStruct) utils.MyStruct {
	if utils.VerificaDocumento(item.Cpf, "cpf") {
		item.Cpf = utils.LimpaCampo(item.Cpf, `\W`)
		item.FlagCPF = "Válido"
	}
	if utils.VerificaDocumento(item.LojaMaisFrequente, "cnpj") {
		item.FlagCNPJFrequente = utils.LimpaCampo(item.FlagCNPJFrequente, `\W`)
		item.FlagCNPJFrequente = "Válido"
	}
	if utils.VerificaDocumento(item.LojaUltimaCompra, "cnpj") {
		item.FlagCNPJUltima = utils.LimpaCampo(item.FlagCNPJUltima, `\W`)
		item.FlagCNPJUltima = "Válido"
	}
	return item
}

// CreateStruct é uma função que cria um array de MyStruct.
// @param fileTextLines: linhas do arquivo que será processado.
func CreateStruct(fileTextLines []string) []utils.MyStruct {
	var objs []utils.MyStruct

	for _, line := range fileTextLines {
		fields := strings.Fields(line)
		if fields[0] == "CPF" {
			continue
		}

		item := utils.MyStruct{
			Cpf:                fields[0],
			Private:            fields[1],
			Incompleto:         fields[2],
			DataUltimaCompra:   fields[3],
			TicketMedio:        fields[4],
			TicketUltimaCompra: fields[5],
			LojaMaisFrequente:  fields[6],
			LojaUltimaCompra:   fields[7],
			FlagCPF:            "Inválido",
			FlagCNPJFrequente:  "Inválido",
			FlagCNPJUltima:     "Inválido",
		}

		newItem := VerificaDocs(item)

		objs = append(objs, newItem)

	}
	return objs
}
