package csvwrapper

import (
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

// OpenFile é uma função que lê um arquivo e retorna o arquivo em si.
func OpenFile(path string) *os.File {
	csvfile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return csvfile
}

// FilterFile é uma função que filtra um dataframe de acordo com uma regra.
func FilterFile(path string) {
	csvfile := OpenFile(path)
	df := dataframe.ReadCSV(csvfile)

	fil := df.Filter(dataframe.F{
		Colname:    "age",
		Comparator: ">",
		Comparando: 20,
	})

	SaveAsCsv(fil)
}

// SaveAsCsv é uma função que salva um dataframe no formato CSV.
func SaveAsCsv(df dataframe.DataFrame) {
	f, err := os.Create("data/output.csv")
	if err != nil {
		log.Fatal(err)
	}
	df.WriteCSV(f)
}
