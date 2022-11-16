package main

import (
	"checkpoint/internal/tickets"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {

	list, err := tickets.SortingByDestination("China")
	if err != nil {
		panic("Couldn't load tickets")
	}

	file, err := os.Create("pessoasPorPaís.csv")
	defer file.Close()
	if err != nil {
		panic("Failed to create file")
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	var data [][]string
	for _, output := range list {
		row := []string{output.Id, output.Name, output.Email, output.Country, output.Time, strconv.Itoa(int(output.Price))}
		data = append(data, row)
	}
	w.WriteAll(data)
	fmt.Println("Relatório gerado, quantidade de viagens por país:", tickets.LineCount("pessoasPorPaís.csv"))

	listOwl, err := tickets.SortingByOwlTime("madrugada")
	if err != nil {
		panic("Couldn't load tickets")
	}

	file1, err := os.Create("viagensMadrugada.csv")
	defer file1.Close()
	if err != nil {
		panic("Failed to create file")
	}

	w1 := csv.NewWriter(file1)
	defer w.Flush()

	var data1 [][]string
	for _, output := range listOwl {
		row1 := []string{output.Id, output.Name, output.Email, output.Country, output.Time, strconv.Itoa(int(output.Price))}
		data1 = append(data1, row1)
	}
	w1.WriteAll(data1)
	fmt.Println("Relatório gerado, quantidade de viagens de madrugada:", tickets.LineCount("viagensMadrugada.csv"))

	listMorning, err := tickets.SortingByMorning("manhã")
	if err != nil {
		panic("Couldn't load tickets")
	}

	file2, err := os.Create("viagensManhã.csv")
	defer file2.Close()
	if err != nil {
		panic("Failed to create file")
	}

	w2 := csv.NewWriter(file2)
	defer w.Flush()

	var data2 [][]string
	for _, output := range listMorning {
		row2 := []string{output.Id, output.Name, output.Email, output.Country, output.Time, strconv.Itoa(int(output.Price))}
		data2 = append(data2, row2)
	}
	w2.WriteAll(data2)
	fmt.Println("Relatório gerado, quantidade de viagens de manhã:", tickets.LineCount("viagensManhã.csv"))

	listAfternoon, err := tickets.SortingByAfternoon("tarde")
	if err != nil {
		panic("Couldn't load tickets")
	}

	file3, err := os.Create("viagensTarde.csv")
	defer file3.Close()
	if err != nil {
		panic("Failed to create file")
	}

	w3 := csv.NewWriter(file3)
	defer w.Flush()

	var data3 [][]string
	for _, output := range listAfternoon {
		row3 := []string{output.Id, output.Name, output.Email, output.Country, output.Time, strconv.Itoa(int(output.Price))}
		data3 = append(data3, row3)
	}
	w3.WriteAll(data3)
	fmt.Println("Relatório gerado, quantidade de viagens de tarde:", tickets.LineCount("viagensTarde.csv"))

	listNight, err := tickets.SortingByNight("noite")
	if err != nil {
		panic("Couldn't load tickets")
	}

	file4, err := os.Create("viagensNoite.csv")
	defer file4.Close()
	if err != nil {
		panic("Failed to create file")
	}

	w4 := csv.NewWriter(file4)
	defer w.Flush()

	var data4 [][]string
	for _, output := range listNight {
		row4 := []string{output.Id, output.Name, output.Email, output.Country, output.Time, strconv.Itoa(int(output.Price))}
		data4 = append(data4, row4)
	}
	w4.WriteAll(data4)
	fmt.Println("Relatório gerado, quantidade de viagens de noite:", tickets.LineCount("viagensNoite.csv"))

	getAverage, err := tickets.AverageDestination("Russia")
	if err != nil {
		panic("Couldn't load average")
	}

	output := fmt.Sprintf("%.2f", getAverage)
	fmt.Println("Média de viagens do destino:", output)
}
