package csvwrapper

import (
	"encoding/csv"
	"fmt"
	"os"
)

type exampleData struct {
	Name string
	Age  string
	City string
}

// OpenFile é uma função que lê um arquivo CSV, com o caminho passado por parâmetro.
func OpenFile(path string) {
	csvFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		emp := exampleData{
			Name: line[0],
			Age:  line[1],
			City: line[2],
		}
		fmt.Println(emp.Name + " " + emp.Age + " " + emp.City)
	}
}
