package tickets

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	Id      string
	Name    string
	Email   string
	Country string
	Time    string
	Price   float64
}

func ReadingFile(path string) ([]Ticket, error) {
	var ticketList []Ticket

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []Ticket{}, err
		}
		ticketList = append(ticketList, Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}
	return ticketList, nil
}

func SortingByDestination(destination string) ([]Ticket, error) {

	ticketList, err := ReadingFile("tickets.csv")
	if err != nil {
		panic("Could not read file.")
	}

	var dest []Ticket

	for _, t := range ticketList {
		if t.Country == destination {
			dest = append(dest, t)
		}
	}

	return dest, nil
}

func SortingByOwlTime(time string) ([]Ticket, error) {

	ticketList, err := ReadingFile("tickets.csv")
	if err != nil {
		panic("Could not read file.")
	}

	var owlTime []Ticket

	for _, t := range ticketList {
		if strings.HasPrefix(t.Time, "0:") || strings.HasPrefix(t.Time, "1:") || strings.HasPrefix(t.Time, "2:") || strings.HasPrefix(t.Time, "3:") || strings.HasPrefix(t.Time, "4:") || strings.HasPrefix(t.Time, "5:") || strings.HasPrefix(t.Time, "6:") && time == "madrugada" {
			owlTime = append(owlTime, t)
		}
	}
	return owlTime, nil
}

func SortingByMorning(time1 string) ([]Ticket, error) {

	morningList, err := ReadingFile("tickets.csv")
	if err != nil {
		panic("Could not read file.")
	}

	var morning []Ticket

	for _, e := range morningList {
		if strings.HasPrefix(e.Time, "7:") || strings.HasPrefix(e.Time, "8:") || strings.HasPrefix(e.Time, "9:") || strings.HasPrefix(e.Time, "10:") || strings.HasPrefix(e.Time, "11:") || strings.HasPrefix(e.Time, "12:") && time1 == "manhã" {
			morning = append(morning, e)
		}
	}
	return morning, nil
}

func SortingByAfternoon(time2 string) ([]Ticket, error) {

	afternoonList, err := ReadingFile("tickets.csv")
	if err != nil {
		panic("Could not read file.")
	}

	var afternoon []Ticket

	for _, f := range afternoonList {
		if strings.HasPrefix(f.Time, "13:") || strings.HasPrefix(f.Time, "14:") || strings.HasPrefix(f.Time, "15:") || strings.HasPrefix(f.Time, "15:") || strings.HasPrefix(f.Time, "16:") || strings.HasPrefix(f.Time, "17:") || strings.HasPrefix(f.Time, "18:") || strings.HasPrefix(f.Time, "19:") && time2 == "tarde" {
			afternoon = append(afternoon, f)
		}
	}
	return afternoon, nil
}

func SortingByNight(time3 string) ([]Ticket, error) {

	nightList, err := ReadingFile("tickets.csv")
	if err != nil {
		panic("Could not read file.")
	}

	var night []Ticket

	for _, g := range nightList {
		if strings.HasPrefix(g.Time, "20:") || strings.HasPrefix(g.Time, "21:") || strings.HasPrefix(g.Time, "22:") || strings.HasPrefix(g.Time, "23:") && time3 == "noite" {
			night = append(night, g)
		}
	}
	return night, nil
}

func CountryMapping(input []Ticket) int {

	mapping := make(map[string]bool)

	for _, ticket := range input {
		mapping[ticket.Country] = true
	}

	return len(mapping)

}

func AverageDestination(destination string) (float64, error) {

	counter := 0

	tickets, err := ReadingFile("tickets.csv")
	if err != nil {
		panic("Could not read file")
	}

	for _, ticket := range tickets {
		if ticket.Country == destination {
			counter++
		}
	}

	return float64(counter) / float64(CountryMapping(tickets)), nil
}

func LineCount(filename string) int64 {
	lc := int64(0)
	f, err := os.Open(filename)
	if err != nil {
		panic("Could not read file")
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		lc++
	}
	return lc
}
