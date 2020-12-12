package export

import (
	"encoding/csv"
	"os"
)

func CreateSheet(terms []string, definitions []string, filetype string, name string) {
	switch filetype {
	case "csv":
		createCsvTsv(terms, definitions, "csv", name)
		break
	case "tsv":
		createCsvTsv(terms, definitions, "tsv", name)
		break
	}

}

func createCsvTsv(terms []string, definitions []string, filetype string, name string) {
	file, err := os.Create(name + "." + filetype)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	if filetype == "tsv" {
		writer.Comma = '\t'
	}
	defer writer.Flush()

	for i := 0; i < len(terms); i++ {
		err := writer.Write([]string{terms[i], definitions[i]})
		if err != nil {
			panic(err)
		}
	}
}
