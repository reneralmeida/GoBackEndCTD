package main

import "fmt"

func main() {
	animalRation, err := animal("spider")
	if err != nil {
		panic(err)
	}
	fmt.Println("Quantidade de ração:", animalRation)
}

func animal(species string) (int, error) {
	base := 1000
	output := base

	switch species {
	case "dog":
		output = base * 10
		return output, nil
	case "cat":
		output = base * 5
		return output, nil
	case "hamster":
		output = base * 250 / 1000
		return output, nil
	case "spider":
		output = base * 150 / 1000
		return output, nil
	}

	return base, fmt.Errorf("Erro ao calcular")
}
