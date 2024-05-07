package csv

import (
	"encoding/csv"
	"os"
)

var ReadAll = func(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := csv.NewReader(file).ReadAll()

	return data
}
