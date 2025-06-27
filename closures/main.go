package main

import "fmt"

// adder returns a closure that adds n to its argument
func adder(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}

func main() {
	add5 := adder(5)
	add10 := adder(10)

	fmt.Println(add5(3))  // 8
	fmt.Println(add10(3)) // 13

	// Closure capturing variable
	setupCounter := func() func() int {
		count := 0
		fmt.Println("Inicializando count com 0")

		// Criando a função interna
		incrementCounter := func() int {
			count++
			fmt.Println("Incrementando count para", count)
			return count
		}

		return incrementCounter
	}

	print("Configurando o contador...")
	// Executando a função externa para obter a função interna
	counter := setupCounter()

	// Agora chamando a função interna múltiplas vezes
	fmt.Println(counter())
	fmt.Println(counter())
}
