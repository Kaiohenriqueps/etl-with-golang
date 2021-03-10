package utils

import "regexp"

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
}

// ChunkSlice é um método que divide um array em array menores.
func ChunkSlice(slice []MyStruct, chunkSize int) [][]MyStruct {
	var chunks [][]MyStruct
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// VerificaCpf é uma função que verifica se o CPF é válido.
func VerificaCpf(cpf string) bool {
	match, _ := regexp.MatchString(`\d{3}\.\d{3}\.\d{3}\-\d{2}`, cpf)
	return match
}
