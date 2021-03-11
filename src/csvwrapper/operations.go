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

// LimpaCampos é uma função que faz uma higienização básica do campo.
func LimpaCampos(item utils.MyStruct) utils.MyStruct {
	item.Cpf = utils.LimpaCampo(item.Cpf, `\W`)
	item.LojaMaisFrequente = utils.LimpaCampo(item.LojaMaisFrequente, `\W`)
	item.LojaUltimaCompra = utils.LimpaCampo(item.LojaUltimaCompra, `\W`)
	return item
}

// VerificaDocs é uma função que verifica se os documentos são válidos.
// @param item: struct do registro.
func VerificaDocs(item utils.MyStruct) utils.MyStruct {
	if utils.VerificaDocumento(item.Cpf, "cpf") {
		item.FlagCPF = "Válido"
	}
	if utils.VerificaDocumento(item.LojaMaisFrequente, "cnpj") {
		item.FlagCNPJFrequente = "Válido"
	}
	if utils.VerificaDocumento(item.LojaUltimaCompra, "cnpj") {
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
			DataUltimaCompra:   utils.VerificaValor(fields[3], "data"),
			TicketMedio:        utils.PreparaNumerico(utils.VerificaValor(fields[4], "numerico")),
			TicketUltimaCompra: utils.PreparaNumerico(utils.VerificaValor(fields[5], "numerico")),
			LojaMaisFrequente:  fields[6],
			LojaUltimaCompra:   fields[7],
			FlagCPF:            "Inválido",
			FlagCNPJFrequente:  "Inválido",
			FlagCNPJUltima:     "Inválido",
		}
		item = LimpaCampos(VerificaDocs(item))
		objs = append(objs, item)

	}
	return objs
}
