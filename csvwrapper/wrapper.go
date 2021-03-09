package csvwrapper

import (
	"fmt"
	"os"
	"github.com/kniren/gota/dataframe"
)


type exampleta struct {
	Name string
	Age string
	City string
}


// OpenFile é uma função queê um arquivo CSV, com o caminho passado por parâmetro.
func OpenFile(path string) {
	csvFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	df := dataframe.ReadCSV(csvFile)
	fmt.Println(df)
	fil := df.Filter(dataframe.F{
		Colname:    "age",
		Comparator: ">",
		Comparando: 20,
	})
	fmt.Println(fil)
}
