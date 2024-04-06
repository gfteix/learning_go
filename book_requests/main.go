package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Person struct {
	BirthYear int    `json:"birth_year"`
	DeathYear int    `json:"death_year"`
	Name      string `json:"name"`
}

type Book struct {
	ID            int               `json:"id"`
	Title         string            `json:"title"`
	Subjects      []string          `json:"subjects"`
	Authors       []Person          `json:"authors"`
	Translators   []Person          `json:"translators"`
	Bookshelves   []string          `json:"bookshelves"`
	Languages     []string          `json:"languages"`
	Copyright     bool              `json:"copyright"`
	MediaType     string            `json:"media_type"`
	Formats       map[string]string `json:"formats"`
	DownloadCount int               `json:"download_count"`
}

type SearchResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Book `json:"results"`
}

func main() {
	fmt.Println("Enter search term:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	search := scanner.Text()

	resp, err := http.Get(fmt.Sprintf("https://gutendex.com/books?search=%s", search))
	if err != nil {
		log.Fatalf("Failed to perform GET request: %v", err)
	}
	defer resp.Body.Close()

	var result SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("Failed to decode response body: %v", err)
	}

	for _, book := range result.Results {
		book.printFields()
	}
}

func (b Book) printFields() {
	fmt.Println("ID:", b.ID)
	fmt.Println("Title:", b.Title)
	fmt.Println("Subjects:", b.Subjects)
	fmt.Println("Authors:", b.Authors)
	fmt.Println("Translators:", b.Translators)
	fmt.Println("Bookshelves:", b.Bookshelves)
	fmt.Println("Languages:", b.Languages)
	fmt.Println("Copyright:", b.Copyright)
	fmt.Println("Media Type:", b.MediaType)
	fmt.Println("Formats:", b.Formats)
	fmt.Println("Download Count:", b.DownloadCount)
	fmt.Println("----------------------------------")
}
