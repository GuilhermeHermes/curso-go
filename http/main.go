package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Post represents a blog post from the JSONPlaceholder API
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	fmt.Println("Starting HTTP client example...")

	// Make a GET request to the JSONPlaceholder API
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return
	}

	// Se você não usar 'defer response.Body.Close()' aqui, acontecerá:
	// 1. Vazamento de recursos - o corpo da resposta não será fechado
	// 2. Conexões TCP podem permanecer abertas
	// 3. Descritores de arquivo não serão liberados
	// 4. Em múltiplas requisições, pode esgotar recursos do sistema
	// 5. Pode causar lentidão e até falhas no programa após uso prolongado

	// Check if the response was successful
	if response.StatusCode != http.StatusOK {
		fmt.Printf("API returned non-200 response: %d %s\n",
			response.StatusCode, response.Status)
		return
	}

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}

	// Parse the JSON response
	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Printf("Error parsing JSON: %s\n", err)
		return
	}

	// Print the post details
	fmt.Println("Successfully fetched post:")
	fmt.Printf("ID: %d\n", post.ID)
	fmt.Printf("User ID: %d\n", post.UserID)
	fmt.Printf("Title: %s\n", post.Title)
	fmt.Printf("Body: %s\n", post.Body)
}
