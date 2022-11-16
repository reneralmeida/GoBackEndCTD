package main

import (
	"fmt"
	"math"
)

type Matriz struct {
	values []float64
	height int
	width  int
}

func (s *Matriz) Set(values []float64) {
	if len(values) != s.height*s.width {
		fmt.Println("valores setados incorreto, matriz invalida")
		return
	}

	s.values = values
}

func (s *Matriz) Print() {
	for row := 0; row < s.height; row++ {
		startAt := s.width * row       // 3*0 = 0 // 3*1 = 3 // 3*2 = 6
		endAt := s.width*row + s.width // 3*0+3 = 3 // 3*1+3 = 6 // 3*2+3 = 9
		fmt.Println(s.values[startAt:endAt])
	}
}

func (s *Matriz) Max() float64 {
	max := -math.MaxFloat64

	for _, v := range s.values {
		if v > max {
			max = v
		}
	}

	return max
}

func (s *Matriz) Quadratic() bool {
	if s.height == s.width && s.height != 0 {
		return true
	}

	return false
}

func main() {
	matriz := Matriz{
		height: 3,
		width:  3,
	}

	matriz.Set([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	matriz.Print()
	fmt.Println("Valor maximo:", matriz.Max())
	fmt.Println("É quadratica:", matriz.Quadratic())
}
