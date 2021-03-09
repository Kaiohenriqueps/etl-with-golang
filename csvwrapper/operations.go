package csvwrapper

import (
	"etl-with-golang/sqlwrapper"
	"fmt"
	"log"
	"os"
	"strconv"

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

	rows := fil.Nrow()
	conn := sqlwrapper.ConnectToPostgres()
	for i := 0; i < rows; i++ {
		name := fil.Elem(i, 0).String()
		age, err := strconv.Atoi(fil.Elem(i, 1).String())
		if err != nil {
			log.Fatal(err)
		}
		city := fil.Elem(i, 2).String()
		values := fmt.Sprintf("'%s', %d, '%s'", name, age, city)
		sqlwrapper.InsertIntoTable(conn, "emp", values)
	}
	// SaveAsCsv(fil)
}

// SaveAsCsv é uma função que salva um dataframe no formato CSV.
func SaveAsCsv(df dataframe.DataFrame) {
	f, err := os.Create("data/output/output.csv")
	if err != nil {
		log.Fatal(err)
	}
	df.WriteCSV(f)
}
