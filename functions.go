package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

var ReadAll = func(fileName string, model any) []any {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := csv.NewReader(file).ReadAll()
	log.Println(data)
	if err != nil {
		for _, row := range data {
			for _, col := range row {
				fmt.Printf("%s,", col)
			}
			fmt.Println()
		}
	}
	return []any{}
}
