package export

import (
	"os"
	"encoding/csv"
)

func CreateSheet(terms []string, definitions []string, filetype string) {
	switch filetype {
	case "csv":
		createCsvTsv(terms, definitions, "csv")
		break;
	case "tsv":
		createCsvTsv(terms, definitions, "tsv")
		break;
	}
	
}

func createCsvTsv(terms []string, definitions []string, filetype string) {
	file, err := os.Create("out." + filetype)
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