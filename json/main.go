package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Pessoa representa uma estrutura básica para uma pessoa
type Pessoa struct {
	Nome      string    `json:"nome"`
	Sobrenome string    `json:"sobrenome"`
	Idade     int       `json:"idade"`
	Email     string    `json:"email,omitempty"` // omitempty omite o campo se for vazio
	Ativo     bool      `json:"ativo"`
	Criado    time.Time `json:"criado"`
}

// NomeCompleto retorna o nome completo da pessoa
func (p Pessoa) NomeCompleto() string {
	return p.Nome + " " + p.Sobrenome
}

// Endereço representa o endereço de uma pessoa
type Endereco struct {
	Rua       string `json:"rua"`
	Numero    int    `json:"numero"`
	Cidade    string `json:"cidade"`
	Estado    string `json:"estado"`
	CEP       string `json:"cep"`
	Pais      string `json:"pais"`
	Principal bool   `json:"principal"`
}

// Usuário representa um usuário com dados aninhados
type Usuario struct {
	ID          int                    `json:"id"`
	Nome        string                 `json:"nome"`
	Email       string                 `json:"email"`
	Senha       string                 `json:"-"` // O hífen omite este campo do JSON
	Enderecos   []Endereco             `json:"enderecos"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	UltimoLogin *time.Time             `json:"ultimo_login,omitempty"`
}

func main() {
	fmt.Println("=== Manipulação de JSON em Go ===")

	// 1. Marshal - Convertendo objeto Go para JSON
	fmt.Println("\n1. Marshal - Convertendo objeto Go para JSON")

	pessoa := Pessoa{
		Nome:      "João",
		Sobrenome: "Silva",
		Idade:     30,
		Email:     "joao@exemplo.com",
		Ativo:     true,
		Criado:    time.Now(),
	}

	// Marshal converte estrutura Go para JSON
	pessoaJSON, err := json.Marshal(pessoa)
	if err != nil {
		log.Fatalf("Erro ao converter para JSON: %v", err)
	}

	// Exibe o JSON resultante
	fmt.Println("JSON gerado:")
	fmt.Println(string(pessoaJSON))

	// 2. MarshalIndent - Gerando JSON formatado
	fmt.Println("\n2. MarshalIndent - Gerando JSON formatado")

	// MarshalIndent gera JSON formatado com indentação
	pessoaJSONFormatado, err := json.MarshalIndent(pessoa, "", "  ")
	if err != nil {
		log.Fatalf("Erro ao converter para JSON formatado: %v", err)
	}

	fmt.Println("JSON formatado:")
	fmt.Println(string(pessoaJSONFormatado))

	// 3. Unmarshal - Convertendo JSON para objeto Go
	fmt.Println("\n3. Unmarshal - Convertendo JSON para objeto Go")

	jsonString := `{
		"nome": "Maria",
		"sobrenome": "Oliveira",
		"idade": 25,
		"email": "maria@exemplo.com",
		"ativo": true,
		"criado": "2023-01-15T10:30:45Z"
	}`

	var novaPessoa Pessoa

	// Unmarshal converte JSON para estrutura Go
	err = json.Unmarshal([]byte(jsonString), &novaPessoa)
	if err != nil {
		log.Fatalf("Erro ao converter JSON para objeto: %v", err)
	}

	fmt.Println("Objeto reconstruído do JSON:")
	fmt.Printf("Nome completo: %s\n", novaPessoa.NomeCompleto())
	fmt.Printf("Idade: %d\n", novaPessoa.Idade)
	fmt.Printf("Email: %s\n", novaPessoa.Email)
	fmt.Printf("Data de criação: %v\n", novaPessoa.Criado)

	// 4. Trabalhando com estruturas aninhadas
	fmt.Println("\n4. Trabalhando com estruturas aninhadas")

	// Criando um usuário com endereços
	usuario := Usuario{
		ID:    1,
		Nome:  "Carlos Pereira",
		Email: "carlos@exemplo.com",
		Senha: "senhasecreta123", // Não será incluída no JSON
		Enderecos: []Endereco{
			{
				Rua:       "Rua das Flores",
				Numero:    123,
				Cidade:    "São Paulo",
				Estado:    "SP",
				CEP:       "01234-567",
				Pais:      "Brasil",
				Principal: true,
			},
			{
				Rua:       "Avenida Principal",
				Numero:    456,
				Cidade:    "Rio de Janeiro",
				Estado:    "RJ",
				CEP:       "21000-000",
				Pais:      "Brasil",
				Principal: false,
			},
		},
		Metadata: map[string]interface{}{
			"ultimo_acesso": "2023-05-20",
			"dispositivo":   "smartphone",
			"versao_app":    2.1,
			"preferencias":  []string{"dark_mode", "notificacoes"},
		},
	}

	// Convertendo para JSON
	usuarioJSON, err := json.MarshalIndent(usuario, "", "  ")
	if err != nil {
		log.Fatalf("Erro ao converter usuário para JSON: %v", err)
	}

	fmt.Println("JSON de usuário com estruturas aninhadas:")
	fmt.Println(string(usuarioJSON))

	// 5. Unmarshal com estruturas aninhadas
	fmt.Println("\n5. Unmarshal com estruturas aninhadas")

	jsonUsuario := `{
		"id": 2,
		"nome": "Ana Santos",
		"email": "ana@exemplo.com",
		"enderecos": [
			{
				"rua": "Rua da Paz",
				"numero": 789,
				"cidade": "Belo Horizonte",
				"estado": "MG",
				"cep": "30000-000",
				"pais": "Brasil",
				"principal": true
			}
		],
		"metadata": {
			"plano": "premium",
			"visitas": 42,
			"interesses": ["tecnologia", "viagem"]
		},
		"ultimo_login": "2023-06-10T15:45:30Z"
	}`

	var novoUsuario Usuario
	err = json.Unmarshal([]byte(jsonUsuario), &novoUsuario)
	if err != nil {
		log.Fatalf("Erro ao converter JSON para usuário: %v", err)
	}

	fmt.Println("Usuário reconstruído do JSON:")
	fmt.Printf("ID: %d\n", novoUsuario.ID)
	fmt.Printf("Nome: %s\n", novoUsuario.Nome)
	fmt.Printf("Email: %s\n", novoUsuario.Email)
	fmt.Printf("Senha (não vem do JSON): %s\n", novoUsuario.Senha)
	fmt.Printf("Quantidade de endereços: %d\n", len(novoUsuario.Enderecos))

	if len(novoUsuario.Enderecos) > 0 {
		endereco := novoUsuario.Enderecos[0]
		fmt.Printf("Primeiro endereço: %s, %d - %s/%s\n",
			endereco.Rua, endereco.Numero, endereco.Cidade, endereco.Estado)
	}

	// Acessando o mapa metadata
	if plano, ok := novoUsuario.Metadata["plano"].(string); ok {
		fmt.Printf("Plano: %s\n", plano)
	}

	if visitas, ok := novoUsuario.Metadata["visitas"].(float64); ok {
		fmt.Printf("Visitas: %.0f\n", visitas)
	}

	if interesses, ok := novoUsuario.Metadata["interesses"].([]interface{}); ok {
		fmt.Printf("Interesses: %v\n", interesses)
	}

	if novoUsuario.UltimoLogin != nil {
		fmt.Printf("Último login: %v\n", *novoUsuario.UltimoLogin)
	}

	// 6. Trabalhando com slices de objetos
	fmt.Println("\n6. Trabalhando com slices de objetos")

	jsonPessoas := `[
		{"nome": "Pedro", "sobrenome": "Costa", "idade": 40, "ativo": true},
		{"nome": "Julia", "sobrenome": "Lima", "idade": 35, "ativo": true},
		{"nome": "Rafael", "sobrenome": "Martins", "idade": 28, "ativo": false}
	]`

	var pessoas []Pessoa
	err = json.Unmarshal([]byte(jsonPessoas), &pessoas)
	if err != nil {
		log.Fatalf("Erro ao converter JSON para slice de pessoas: %v", err)
	}

	fmt.Println("Lista de pessoas do JSON:")
	for i, p := range pessoas {
		fmt.Printf("%d: %s %s, %d anos, ativo: %t\n",
			i+1, p.Nome, p.Sobrenome, p.Idade, p.Ativo)
	}

	// 7. Trabalhando com mapas
	fmt.Println("\n7. Trabalhando com mapas")

	jsonMapa := `{
		"BR": "Brasil",
		"US": "Estados Unidos",
		"FR": "França",
		"JP": "Japão",
		"AU": "Austrália"
	}`

	// Declarando um mapa para receber os dados
	paises := make(map[string]string)

	err = json.Unmarshal([]byte(jsonMapa), &paises)
	if err != nil {
		log.Fatalf("Erro ao converter JSON para mapa: %v", err)
	}

	fmt.Println("Mapa de países:")
	for codigo, nome := range paises {
		fmt.Printf("%s: %s\n", codigo, nome)
	}

	// 8. Decodificando JSON de streams
	fmt.Println("\n8. Decodificando JSON de streams")

	jsonStream := strings.NewReader(`{"nome": "Lucas", "sobrenome": "Ferreira", "idade": 32}`)

	// Criando um decoder para ler do stream
	decoder := json.NewDecoder(jsonStream)

	var pessoaStream Pessoa
	err = decoder.Decode(&pessoaStream)
	if err != nil {
		log.Fatalf("Erro ao decodificar JSON do stream: %v", err)
	}

	fmt.Println("Pessoa decodificada do stream:")
	fmt.Printf("%s %s, %d anos\n", pessoaStream.Nome, pessoaStream.Sobrenome, pessoaStream.Idade)

	// 9. Encoding JSON para streams
	fmt.Println("\n9. Encoding JSON para streams")

	// Criando um buffer para simular um writer
	var jsonBuffer strings.Builder

	// Criando um encoder para escrever no buffer
	encoder := json.NewEncoder(&jsonBuffer)

	// Configurando indentação (opcional)
	encoder.SetIndent("", "  ")

	// Objeto para codificar
	pessoaParaStream := Pessoa{
		Nome:      "Fernanda",
		Sobrenome: "Souza",
		Idade:     29,
		Email:     "fernanda@exemplo.com",
		Ativo:     true,
		Criado:    time.Now(),
	}

	// Codificando o objeto para o buffer
	err = encoder.Encode(pessoaParaStream)
	if err != nil {
		log.Fatalf("Erro ao codificar objeto para JSON: %v", err)
	}

	fmt.Println("JSON codificado para o stream:")
	fmt.Println(jsonBuffer.String())

	// 10. Salvando JSON em arquivo
	fmt.Println("\n10. Salvando JSON em arquivo")

	// Lista de itens para salvar
	itens := []map[string]interface{}{
		{
			"id":         1,
			"nome":       "Laptop",
			"preco":      3500.99,
			"disponivel": true,
			"categorias": []string{"eletrônicos", "computadores"},
		},
		{
			"id":         2,
			"nome":       "Smartphone",
			"preco":      2100.50,
			"disponivel": true,
			"categorias": []string{"eletrônicos", "celulares"},
		},
		{
			"id":         3,
			"nome":       "Headphone",
			"preco":      450.75,
			"disponivel": false,
			"categorias": []string{"eletrônicos", "acessórios"},
		},
	}

	// Convertendo para JSON formatado
	itensJSON, err := json.MarshalIndent(itens, "", "  ")
	if err != nil {
		log.Fatalf("Erro ao converter itens para JSON: %v", err)
	}

	// Salvando em arquivo
	err = os.WriteFile("itens.json", itensJSON, 0644)
	if err != nil {
		log.Fatalf("Erro ao salvar JSON em arquivo: %v", err)
	}

	fmt.Println("JSON salvo com sucesso no arquivo 'itens.json'")

	// 11. Lendo JSON de arquivo
	fmt.Println("\n11. Lendo JSON de arquivo")

	// Verificando se o arquivo existe (acabamos de criá-lo)
	jsonBytes, err := os.ReadFile("itens.json")
	if err != nil {
		log.Fatalf("Erro ao ler arquivo JSON: %v", err)
	}

	var itensLidos []map[string]interface{}
	err = json.Unmarshal(jsonBytes, &itensLidos)
	if err != nil {
		log.Fatalf("Erro ao converter JSON para objetos: %v", err)
	}

	fmt.Printf("Lidos %d itens do arquivo:\n", len(itensLidos))
	for i, item := range itensLidos {
		fmt.Printf("Item %d: %s - R$ %.2f\n",
			i+1, item["nome"], item["preco"])
	}

	fmt.Println("\n=== Fim do Exemplo ===")
}

// Um tipo customizado para demonstrar Marshal/Unmarshal personalizados
type DataPersonalizada struct {
	Dia int
	Mes int
	Ano int
}

// MarshalJSON implementa a interface Marshaler
func (d DataPersonalizada) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%02d/%02d/%04d\"", d.Dia, d.Mes, d.Ano)), nil
}

// UnmarshalJSON implementa a interface Unmarshaler
func (d *DataPersonalizada) UnmarshalJSON(data []byte) error {
	var dataStr string
	if err := json.Unmarshal(data, &dataStr); err != nil {
		return err
	}

	_, err := fmt.Sscanf(dataStr, "%d/%d/%d", &d.Dia, &d.Mes, &d.Ano)
	return err
}
