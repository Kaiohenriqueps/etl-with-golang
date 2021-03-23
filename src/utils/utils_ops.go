package utils

import (
	"regexp"
	"strings"
)

// MyStruct é a estrutura do item de cada linha.
type MyStruct struct {
	Cpf                string
	Private            string
	Incompleto         string
	DataUltimaCompra   string
	TicketMedio        string
	TicketUltimaCompra string
	LojaMaisFrequente  string
	LojaUltimaCompra   string
	FlagCPF            string
	FlagCNPJFrequente  string
	FlagCNPJUltima     string
}

// Regex é uma variável que contém os regex dos documentos.
var Regex = map[string]string{
	"cpf":  `\d{3}\.\d{3}\.\d{3}\-\d{2}`,
	"cnpj": `\d{2}\.\d{3}\.\d{3}\/\d{4}\-\d{2}`}

// Data padrao.
var valoresPadrao = map[string]string{
	"data":     "1111-01-01",
	"numerico": "0",
	"text":     "NA"}

// VerificaDocumento é uma função que verifica se o CNPJ ou CPF são válidos.
func VerificaDocumento(doc string, docType string) bool {
	match, _ := regexp.MatchString(Regex[docType], doc)
	return match
}

// LimpaCampo é uma função que recebe uma string e retira os caracteres especiais.
func LimpaCampo(str string, regex string) string {
	if str == "NULL" || str == "" {
		return ""
	}
	reg := regexp.MustCompile(regex)
	return reg.ReplaceAllString(str, "")
}

// VerificaValor é uma função que recebe um valor, caso ele seja vazio ou NULL, retorna um valor padrão.
// Caso contrário, retorna o valor do campo.
func VerificaValor(data string, valueType string) string {
	if data == "NULL" || data == "" {
		return valoresPadrao[valueType]
	}
	return data
}

// PreparaNumerico é uma função que prepara a string para a inserção de um numérico.
func PreparaNumerico(numero string) string {
	if strings.Contains(numero, ",") {
		return strings.ReplaceAll(numero, ",", ".")
	}
	return numero
}
