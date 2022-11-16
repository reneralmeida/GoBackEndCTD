package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type ListaProdutos struct {
	ID   int
	Name string
}

type Product1 struct {
	Code       int
	Name       string
	Price      int
	Quantidade int
}

func main() {
	products := []ListaProdutos{
		{01, "Appliances"},
		{02, "Furniture"},
		{03, "Tools"},
		{04, "Whatever"},
	}

	products1 := []Product1{
		{01, "TV LG", 1500, 20},
		{02, "TV Samsung", 2000, 30},
		{03, "Micro System Sony", 1000, 15},
		{04, "Boombox", 250, 10},
	}

	file, err := os.Create("lista.csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	var data [][]string
	for _, product := range products {
		row := []string{strconv.Itoa(product.ID), product.Name}
		data = append(data, row)
	}
	w.WriteAll(data)

	file1, err := os.Create("produtos.csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	e := csv.NewWriter(file1)
	defer w.Flush()

	var data1 [][]string
	for _, produto := range products1 {
		row := []string{strconv.Itoa(produto.Code), produto.Name, strconv.Itoa(produto.Price), strconv.Itoa(produto.Quantidade)}
		data1 = append(data1, row)
	}
	e.WriteAll(data1)
}
