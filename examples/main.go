// example/main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AlexanderXinarxZenDev/mango_template"
)

func main() {
	// Create a sample template file
	content := `<h1>{{ title }}</h1>
<p>{{ message }}</p>
<span>{{ missing }}</span>`
	err := os.WriteFile("index.html", []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Render with data
	data := map[string]any{
		"title":   "Flusk App",
		"message": "Hello MangoTemplate",
	}

	result, err := mango.Render("index.html", data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	// Output:
	// <h1>Flusk App</h1>
	// <p>Hello MangoTemplate</p>
	// <span></span>
}