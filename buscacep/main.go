package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Endereco struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func buscarCep(cep string) (Endereco, error) {
	url := "https://viacep.com.br/ws/" + cep + "/json/"

	response, err := http.Get(url)
	if err != nil {
		return Endereco{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Endereco{}, err
	}

	var endereco Endereco

	err = json.NewDecoder(response.Body).Decode(&endereco)
	if err != nil {
		return Endereco{}, err
	}

	return endereco, nil
}

func main() {
	log.Println("Iniciando servidor na porta 8080...")
	listenAndServeTLS()
}
