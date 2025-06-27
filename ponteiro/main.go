package main

import "fmt"

func incrementa(x *int) {
	*x = *x + 1
}

func trocaValores(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func zeraValor(x *int) {
	*x = 0
}

func main() {
	a := 5
	incrementa(&a)
	fmt.Println("Valor de a após incremento:", a)

	// Exemplo de troca de valores usando ponteiros
	x, y := 10, 20
	fmt.Println("Antes da troca: x =", x, ", y =", y)
	trocaValores(&x, &y)
	fmt.Println("Depois da troca: x =", x, ", y =", y)

	// Exemplo de zerar valor usando ponteiro
	z := 100
	fmt.Println("Antes de zerar: z =", z)
	zeraValor(&z)
	fmt.Println("Depois de zerar: z =", z)
	// Exemplo didático de ponteiro para ponteiro
	fmt.Println("\n=== EXEMPLO DE PONTEIRO PARA PONTEIRO ===\n")

	// Criando uma variável normal com valor 42
	num := 42
	fmt.Println("1. Variável num criada com valor:", num)

	// Criando um ponteiro que aponta para num
	ptr := &num
	fmt.Printf("2. Ponteiro ptr criado apontando para num (endereço %p)\n", ptr)

	// Criando um ponteiro para ponteiro (ptr2 aponta para ptr)
	ptr2 := &ptr
	fmt.Printf("3. Ponteiro para ponteiro ptr2 criado apontando para ptr (endereço %p)\n", ptr2)

	fmt.Println("\n=== ACESSANDO VALORES ===\n")

	// Acessando o valor diretamente
	fmt.Println("4. Valor direto de num:", num)

	// Acessando o valor através do ponteiro (desreferenciação simples)
	fmt.Println("5. Valor de num via ptr (*ptr):", *ptr)

	// Acessando o valor através do ponteiro para ponteiro (desreferenciação dupla)
	fmt.Println("6. Valor de num via ptr2 (**ptr2):", **ptr2)

	fmt.Println("\n=== MODIFICANDO O VALOR ===\n")

	// Modificando o valor da variável original
	num = 100
	fmt.Println("7. Valor de num modificado para:", num)

	// O ponteiro reflete a mudança automaticamente
	fmt.Println("8. Valor via ponteiro após mudança (*ptr):", *ptr)

	// O ponteiro para ponteiro também reflete a mudança
	fmt.Println("9. Valor via ponteiro para ponteiro após mudança (**ptr2):", **ptr2)

	fmt.Println("\n=== VISUALIZANDO ENDEREÇOS ===\n")

	// Diagrama visual da cadeia de referências
	fmt.Println("DIAGRAMA DA CADEIA DE PONTEIROS:")
	fmt.Printf("num (valor: %d) ← ptr (endereço: %p) ← ptr2 (endereço: %p)\n\n", num, ptr, ptr2)

	// Endereços de memória
	fmt.Printf("10. Endereço da variável num (&num): %p\n", &num)
	fmt.Printf("11. Valor armazenado em ptr (ptr): %p\n", ptr)
	fmt.Printf("12. Endereço da variável ptr (&ptr): %p\n", &ptr)
	fmt.Printf("13. Valor armazenado em ptr2 (ptr2): %p\n", ptr2)
	fmt.Printf("14. Endereço da variável ptr2 (&ptr2): %p\n", &ptr2)

	fmt.Println("\n=== DESREFERENCIAÇÃO ===\n")

	fmt.Println("15. *ptr acessa o valor de num:", *ptr)
	fmt.Println("16. *ptr2 acessa o valor de ptr (que é o endereço de num):", *ptr2)
	fmt.Println("17. **ptr2 acessa o valor de num:", **ptr2)

	fmt.Println("\n=== RESUMO ===\n")
	fmt.Println("• num = valor original (100)")
	fmt.Println("• &num = endereço da variável num")
	fmt.Println("• ptr = &num (ponteiro para num)")
	fmt.Println("• *ptr = valor de num (100)")
	fmt.Println("• &ptr = endereço da variável ptr")
	fmt.Println("• ptr2 = &ptr (ponteiro para ptr)")
	fmt.Println("• *ptr2 = valor de ptr (que é &num)")
	fmt.Println("• **ptr2 = valor de num (100)")
}
