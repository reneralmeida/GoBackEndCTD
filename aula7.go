package main

import (
	"errors"
	"fmt"
)

const lowSalary = 330

var ErrSalary = errors.New("salario baixo")

func isLowerThanMin(salary int) error {
	if salary < lowSalary {
		return fmt.Errorf("validade salary: %w", ErrSalary)
	}
	return nil
}

func main() {
	salaryLower := 250
	err := isLowerThanMin(salaryLower)
	if errors.Is(err, ErrSalary) {
		fmt.Println(err, err.Error())
	}
}
