package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type ListaClientes struct {
	Code     int
	Name     string
	Category string
}

func main() {
	customers := []ListaClientes{
		{01, "Joelmir", "VIP"},
		{02, "Joel", "Vitalício"},
		{03, "Tools", "Free-tier"},
		{04, "Whatever", "VIP"},
	}

	file, err := os.Create("lista1.csv")
	defer file.Close()
	if err != nil {
		panic("failed to create file")
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	var data [][]string
	for _, customer := range customers {
		row := []string{strconv.Itoa(customer.Code), customer.Name, customer.Category}
		data = append(data, row)
	}
	w.WriteAll(data)

	file1, err := os.Open("lista1.csv")
	defer file1.Close()
	if err != nil {
		panic("failed to open file")
	}

	readingFile, err := os.ReadFile("lista1.csv")
	if err != nil {
		panic("failed to read file")
	}

	fmt.Printf(string(readingFile))

}
