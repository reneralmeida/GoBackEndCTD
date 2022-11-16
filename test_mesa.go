package main

import "fmt"

func main() {
	salario, err := calculaSalario(3000, "C")
	if err != nil {
		panic(err)
	}
	fmt.Println("Salário:", salario)
}

func calculaSalario(minutos int, categoria string) (float64, error) {
	salario := 0.0
	horas := minutos / 60
	switch categoria {
	case "A":
		salario = (float64(horas) * 112.38) * 1.5
		return salario, nil
	case "B":
		salario = (float64(horas) * 56.19) * 1.2
		return salario, nil
	case "C":
		salario = float64(horas) * 37.46
		return salario, nil
	}

	return salario, fmt.Errorf("Erro ao calcular")
}
