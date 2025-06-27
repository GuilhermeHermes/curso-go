package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// criarArquivo cria um novo arquivo com o nome especificado
func criarArquivo(nome string) error {
	arquivo, err := os.Create(nome)
	if err != nil {
		return err
	}
	defer arquivo.Close()
	return nil
}

// escreverEmArquivo escreve texto em um arquivo (sobrescreve se já existir)
func escreverEmArquivo(nome string, conteudo string) error {
	arquivo, err := os.Create(nome)
	if err != nil {
		return err
	}
	defer arquivo.Close()

	_, err = arquivo.WriteString(conteudo)
	return err
}

// adicionarEmArquivo adiciona texto ao final de um arquivo existente
func adicionarEmArquivo(nome string, conteudo string) error {
	arquivo, err := os.OpenFile(nome, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer arquivo.Close()

	_, err = arquivo.WriteString(conteudo)
	return err
}

func lerArquivoCompleto(nome string) (string, error) {
	conteudo, err := ioutil.ReadFile(nome)
	if err != nil {
		return "", err
	}
	return string(conteudo), nil
}

// lerArquivoLinha lê um arquivo linha por linha
func lerArquivoLinha(nome string) error {
	arquivo, err := os.Open(nome)
	if err != nil {
		return err
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return scanner.Err()
}

// verificarExistencia verifica se um arquivo ou diretório existe
func verificarExistencia(nome string) bool {
	_, err := os.Stat(nome)
	return !os.IsNotExist(err)
}

// obterInformacoes obtém informações sobre um arquivo
func obterInformacoes(nome string) error {
	info, err := os.Stat(nome)
	if err != nil {
		return err
	}

	fmt.Printf("Nome: %s\n", info.Name())
	fmt.Printf("Tamanho: %d bytes\n", info.Size())
	fmt.Printf("Permissões: %s\n", info.Mode())
	fmt.Printf("Última modificação: %s\n", info.ModTime())
	fmt.Printf("É diretório: %t\n", info.IsDir())

	return nil
}

// copiarArquivo copia um arquivo de origem para destino
func copiarArquivo(origem, destino string) error {
	// Abrir arquivo de origem
	original, err := os.Open(origem)
	if err != nil {
		return err
	}
	defer original.Close()

	// Criar arquivo de destino
	novo, err := os.Create(destino)
	if err != nil {
		return err
	}
	defer novo.Close()

	// Copiar conteúdo
	_, err = io.Copy(novo, original)
	return err
}

// renomearArquivo renomeia ou move um arquivo
func renomearArquivo(origem, destino string) error {
	return os.Rename(origem, destino)
}

// excluirArquivo remove um arquivo
func excluirArquivo(nome string) error {
	return os.Remove(nome)
}

// criarDiretorio cria um novo diretório
func criarDiretorio(nome string) error {
	return os.Mkdir(nome, 0755)
}

// listarDiretorio lista os arquivos em um diretório
func listarDiretorio(diretorio string) error {
	arquivos, err := ioutil.ReadDir(diretorio)
	if err != nil {
		return err
	}

	for _, arquivo := range arquivos {
		fmt.Println(arquivo.Name())
	}

	return nil
}

// percorrerDiretorio percorre recursivamente um diretório
func percorrerDiretorio(raiz string) error {
	return filepath.Walk(raiz, func(caminho string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("Caminho: %s, É diretório: %t\n", caminho, info.IsDir())
		return nil
	})
}

// exemploDefer demonstra o funcionamento do defer
func exemploDefer() {
	// Exemplo 1: Garantindo que um arquivo seja fechado
	arquivo, err := os.Open("exemplo.txt")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return
	}
	// Este defer garante que arquivo.Close() será chamado quando a função terminar
	// mesmo que ocorra um erro ou um return antecipado
	defer arquivo.Close()

	// Exemplo 2: Múltiplos defer são executados em ordem LIFO (último a entrar, primeiro a sair)
	defer fmt.Println("1. Esta mensagem aparece por último")
	defer fmt.Println("2. Esta mensagem aparece em segundo")
	defer fmt.Println("3. Esta mensagem aparece primeiro")

	// Exemplo 3: defer com função anônima
	defer func() {
		fmt.Println("4. Executando uma função anônima adiada")
	}()

	// O resto do código da função continua normalmente...
	fmt.Println("Processando o arquivo...")

	// Quando a função terminar, os defer serão executados na ordem inversa
}

func main() {
	// Arquivo para os testes
	arquivoTeste := "exemplo.txt"
	arquivoCopia := "exemplo_copia.txt"
	arquivoRenomeado := "exemplo_renomeado.txt"
	diretorioTeste := "pasta_teste"

	fmt.Println("=== Exemplo de Manipulação de Arquivos em Go ===")

	// Criar um arquivo
	fmt.Println("\n1. Criando arquivo...")
	if err := criarArquivo(arquivoTeste); err != nil {
		fmt.Printf("Erro ao criar arquivo: %v\n", err)
	} else {
		fmt.Println("Arquivo criado com sucesso.")
	}

	// Escrever no arquivo
	fmt.Println("\n2. Escrevendo no arquivo...")
	conteudo := "Linha 1: Olá, Mundo!\nLinha 2: Este é um exemplo de manipulação de arquivos em Go.\n"
	if err := escreverEmArquivo(arquivoTeste, conteudo); err != nil {
		fmt.Printf("Erro ao escrever no arquivo: %v\n", err)
	} else {
		fmt.Println("Conteúdo escrito com sucesso.")
	}

	// Adicionar conteúdo ao arquivo
	fmt.Println("\n3. Adicionando conteúdo ao arquivo...")
	novoConteudo := "Linha 3: Esta linha foi adicionada posteriormente.\n"
	if err := adicionarEmArquivo(arquivoTeste, novoConteudo); err != nil {
		fmt.Printf("Erro ao adicionar conteúdo: %v\n", err)
	} else {
		fmt.Println("Conteúdo adicionado com sucesso.")
	}

	// Ler o arquivo completo
	fmt.Println("\n4. Lendo o arquivo completo...")
	if conteudoLido, err := lerArquivoCompleto(arquivoTeste); err != nil {
		fmt.Printf("Erro ao ler arquivo: %v\n", err)
	} else {
		fmt.Println("Conteúdo do arquivo:")
		fmt.Println("---")
		fmt.Print(conteudoLido)
		fmt.Println("---")
	}

	// Ler o arquivo linha a linha
	fmt.Println("\n5. Lendo o arquivo linha por linha...")
	fmt.Println("Conteúdo do arquivo (linha por linha):")
	fmt.Println("---")
	if err := lerArquivoLinha(arquivoTeste); err != nil {
		fmt.Printf("Erro ao ler arquivo linha por linha: %v\n", err)
	}
	fmt.Println("---")

	// Verificar informações do arquivo
	fmt.Println("\n6. Informações do arquivo:")
	if err := obterInformacoes(arquivoTeste); err != nil {
		fmt.Printf("Erro ao obter informações: %v\n", err)
	}

	// Verificar existência
	fmt.Println("\n7. Verificando existência de arquivos...")
	fmt.Printf("O arquivo '%s' existe? %t\n", arquivoTeste, verificarExistencia(arquivoTeste))
	fmt.Printf("O arquivo 'nao_existe.txt' existe? %t\n", verificarExistencia("nao_existe.txt"))

	// Copiar arquivo
	fmt.Println("\n8. Copiando arquivo...")
	if err := copiarArquivo(arquivoTeste, arquivoCopia); err != nil {
		fmt.Printf("Erro ao copiar arquivo: %v\n", err)
	} else {
		fmt.Println("Arquivo copiado com sucesso.")
		fmt.Printf("A cópia '%s' existe? %t\n", arquivoCopia, verificarExistencia(arquivoCopia))
	}

	// Renomear arquivo
	fmt.Println("\n9. Renomeando arquivo...")
	if err := renomearArquivo(arquivoCopia, arquivoRenomeado); err != nil {
		fmt.Printf("Erro ao renomear arquivo: %v\n", err)
	} else {
		fmt.Println("Arquivo renomeado com sucesso.")
		fmt.Printf("O arquivo '%s' existe? %t\n", arquivoRenomeado, verificarExistencia(arquivoRenomeado))
	}

	// Criar diretório
	fmt.Println("\n10. Criando diretório...")
	if err := criarDiretorio(diretorioTeste); err != nil {
		fmt.Printf("Erro ao criar diretório: %v\n", err)
	} else {
		fmt.Println("Diretório criado com sucesso.")
	}

	// Listar diretório atual
	fmt.Println("\n11. Listando arquivos no diretório atual...")
	diretorioAtual, _ := os.Getwd()
	if err := listarDiretorio(diretorioAtual); err != nil {
		fmt.Printf("Erro ao listar diretório: %v\n", err)
	}

	// Percorrer diretório recursivamente
	fmt.Println("\n12. Percorrendo diretório recursivamente...")
	if err := percorrerDiretorio(diretorioAtual); err != nil {
		fmt.Printf("Erro ao percorrer diretório: %v\n", err)
	}

	// Excluir arquivos
	fmt.Println("\n13. Excluindo arquivos...")
	if err := excluirArquivo(arquivoRenomeado); err != nil {
		fmt.Printf("Erro ao excluir arquivo '%s': %v\n", arquivoRenomeado, err)
	} else {
		fmt.Printf("Arquivo '%s' excluído com sucesso.\n", arquivoRenomeado)
	}

	fmt.Println("\n=== Fim do Exemplo ===")
}
