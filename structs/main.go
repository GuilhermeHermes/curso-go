package main

import (
	"fmt"
)

// Person defines a struct with name and age fields
type Person struct {
	Name  string
	Age   int
	Ativo bool
}

func Ativar(p *Person) {
	p.Ativo = true
	fmt.Println("Pessoa ativada:", p.Name)
	fmt.Println("Ativo status:", p.Ativo)
}

func (p *Person) Desativar() {
	p.Ativo = false
	fmt.Println("Pessoa desativada:", p.Name)
	fmt.Println("Ativo status:", p.Ativo)
}

func main() {
	// Create a new Person instance
	alice := Person{Name: "Alice", Age: 30, Ativo: true}

	alice.Desativar()

	// Access struct fields
	fmt.Printf("Name: %s, Age: %d, Ativo: %t\n", alice.Name, alice.Age, alice.Ativo)

	Ativar(&alice)
	fmt.Printf("Ativo status after activation: %t\n", alice.Ativo)

	// Update struct fields
	alice.Age = 31
	fmt.Printf("Updated Age: %d\n", alice.Age)

	// Create a pointer to a struct
	p2 := &Person{Name: "Bob", Age: 25}
	fmt.Printf("Pointer - Name: %s, Age: %d\n", p2.Name, p2.Age)

	printPessoa()
}
