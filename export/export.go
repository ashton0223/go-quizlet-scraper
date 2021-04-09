package export

import (
	"encoding/csv"
	"errors"
	"os"
)

func CreateSheet(terms []string, definitions []string, filetype string, name string) error {
	switch filetype {
	case "csv":
		return createCsvTsv(terms, definitions, "csv", name)
	case "tsv":
		return createCsvTsv(terms, definitions, "tsv", name)
	}
	return errors.New("Not tsv or csv")
}

func createCsvTsv(terms []string, definitions []string, filetype string, name string) error {
	file, err := os.Create(name + "." + filetype)
	if err != nil {
		return err
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
			return err
		}
	}

	return nil
}
