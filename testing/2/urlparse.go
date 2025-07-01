package urlparser

import (
	"errors"
	"strings"
)

type URLParams struct {
	Protocol    string
	Domain      string
	Path        []string
	QueryParams map[string]string
	Fragment    string
}

func ParseCustomURL(url string) (URLParams, error) {
	result := URLParams{
		QueryParams: make(map[string]string),
	}

	if url == "" {
		return result, errors.New("URL vazia")
	}

	// Separar protocolo
	protocolParts := strings.SplitN(url, "://", 2)
	if len(protocolParts) != 2 {
		return result, errors.New("formato de protocolo inválido")
	}

	result.Protocol = protocolParts[0]
	remainder := protocolParts[1]

	// Separar fragmento
	if strings.Contains(remainder, "#") {
		parts := strings.SplitN(remainder, "#", 2)
		remainder = parts[0]
		result.Fragment = parts[1]
	}

	// Separar query params
	if strings.Contains(remainder, "?") {
		parts := strings.SplitN(remainder, "?", 2)
		remainder = parts[0]
		queryString := parts[1]

		queryParts := strings.Split(queryString, "&")
		for _, part := range queryParts {
			if part == "" {
				continue
			}

			keyValue := strings.SplitN(part, "=", 2)
			if len(keyValue) == 1 {
				// Parâmetro sem valor, armazenar com valor vazio
				result.QueryParams[keyValue[0]] = ""
			} else if len(keyValue) == 2 {
				// Parâmetro com valor normal
				result.QueryParams[keyValue[0]] = keyValue[1]
			}
			// Não retornamos erro para formatos inválidos,
			// apenas processamos o que pudermos
		}
	}

	// Separar domínio e path
	domainAndPath := strings.SplitN(remainder, "/", 2)
	result.Domain = domainAndPath[0]

	if len(domainAndPath) > 1 {
		pathParts := strings.Split(domainAndPath[1], "/")
		for _, part := range pathParts {
			if part != "" {
				result.Path = append(result.Path, part)
			}
		}
	}

	return result, nil
}
