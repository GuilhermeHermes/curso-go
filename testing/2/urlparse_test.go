package urlparser

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func FuzzParseCustomURL(f *testing.F) {
	// Seeds básicos
	f.Add("http://example.com")
	f.Add("https://api.example.org/v1/users?id=123&filter=active#results")
	f.Add("ftp://files.example.net/public/docs/")

	f.Fuzz(func(t *testing.T, url string) {
		// Ignora strings não-UTF8 válidas
		if !utf8.ValidString(url) {
			return
		}

		result, err := ParseCustomURL(url)

		if err == nil {
			// Validações de consistência nos dados obtidos

			// 1. Se temos um protocolo, ele deve terminar com ://
			if strings.Contains(url, "://") {
				protoParts := strings.SplitN(url, "://", 2)
				if result.Protocol != protoParts[0] {
					t.Errorf("Protocolo extraído incorretamente: %s", url)
				}
			}

			// 2. Se temos fragmento (#), deve estar presente no resultado
			if strings.Contains(url, "#") {
				if !strings.Contains(url, "#"+result.Fragment) {
					t.Errorf("Fragmento extraído incorretamente: %s", url)
				}
			}

			// 3. Todos os parâmetros da query devem existir no resultado
			if strings.Contains(url, "?") {
				queryPart := strings.SplitN(url, "?", 2)[1]
				if strings.Contains(queryPart, "#") {
					queryPart = strings.SplitN(queryPart, "#", 2)[0]
				}

				queryParams := strings.Split(queryPart, "&")
				for _, param := range queryParams {
					if param == "" {
						continue
					}

					keyValue := strings.SplitN(param, "=", 2)
					if len(keyValue) == 2 && keyValue[0] != "" {
						if val, exists := result.QueryParams[keyValue[0]]; !exists || val != keyValue[1] {
							t.Errorf("Parâmetro '%s' não extraído corretamente em: %s", keyValue[0], url)
						}
					}
				}
			}

			// 4. Reconstrução do caminho deve corresponder ao original
			if len(result.Path) > 0 {
				reconstructedPath := "/" + strings.Join(result.Path, "/")
				if !strings.Contains(url, reconstructedPath) {
					t.Errorf("Caminho reconstruído incorretamente: %s", url)
				}
			}
		} else {
			// Para URLs com erro, verificamos se realmente deveria falhar
			// Por exemplo, uma URL sem protocolo deve falhar
			if strings.Contains(url, "://") && !strings.HasPrefix(url, "://") && strings.Contains(url, ".") {
				// URLs que parecem válidas não deveriam falhar
				t.Errorf("URL potencialmente válida falhou: %s com erro: %v", url, err)
			}
		}
	})
}
