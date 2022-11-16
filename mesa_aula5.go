package main

import "fmt"

const (
	PEQUENO = "pequeno"
	MEDIO   = "medio"
	GRANDE  = "grande"
)

type Produto struct {
	valor float64
}

func (p Produto) preco() float64 {
	return p.valor
}

type Empresa struct {
	Pequeno Produto
	Medio   Produto
	Grande  Produto
}

func (v *Empresa) factoryProduto(typeProduto string, valor float64) Produto {
	p := Produto{valor: valor}
	switch typeProduto {
	case PEQUENO:
		v.Pequeno = p
		return p
	case MEDIO:
		p.valor = p.valor * 1.03
		v.Medio = p
		return p
	case GRANDE:
		p.valor = (p.valor * 1.06) + 50
		v.Grande = p
		return p
	}
	return Produto{}
}

func (v *Empresa) preco() float64 {
	return (v.Pequeno.preco()) + (v.Grande.preco()) + (v.Medio.preco())
}

func main() {

	dh := Empresa{}
	dh.factoryProduto("pequeno", 2500.0)
	dh.factoryProduto("medio", 3700.0)
	dh.factoryProduto("grande", 35000.0)
	fmt.Println(dh)
	fmt.Println(dh.preco())
}
