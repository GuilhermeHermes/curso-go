package main

import (
	"encoding/json"
	"net/http"
)

func listenAndServeTLS() {
	http.HandleFunc("/", buscaCepHandler)
	http.ListenAndServe(":8080", nil)
}

func buscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Parâmetro 'cep' é obrigatório"))
		return
	}

	endereco, err := buscarCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao buscar CEP: " + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(endereco)
}
