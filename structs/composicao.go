package main

import "fmt"

// Endereco é uma struct que será composta em outra struct
type Endereco struct {
	Rua    string
	Cidade string
	Estado string
}

// Pessoa compõe Endereco
type Pessoa struct {
	Nome     string
	Idade    int
	Endereco // composição: Pessoa "tem um" Endereco
}

func printPessoa() {
	p := Pessoa{
		Nome:  "João",
		Idade: 30,
		Endereco: Endereco{
			Rua:    "Rua das Flores",
			Cidade: "São Paulo",
			Estado: "SP",
		},
	}

	fmt.Printf("Nome: %s\nIdade: %d\nEndereço: %s, %s - %s\n",
		p.Nome, p.Idade, p.Rua, p.Cidade, p.Estado)
}
