package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Catalog struct {
	XMLName xml.Name `xml:"catalog" json:"-"`
	Text    string   `xml:",chardata" json:"-"`
	Book    []struct {
		Text        string `xml:",chardata" json:"-"`
		ID          string `xml:"id,attr" json:"id"`
		Author      string `xml:"author" json:"author"`
		Title       string `xml:"title"  json:"title"`
		Genre       string `xml:"genre"  json:"genre"`
		Price       string `xml:"price" json:"price"`
		PublishDate string `xml:"publish_date" json:"publishDate"`
		Description string `xml:"description"  json:"description"`
	} `xml:"book" json:"books"`
}

// Wrapper struct for JSON output
type CatalogJson struct {
	Catalog Catalog `json:"catalog"`
}

func main() {
	file, err := os.Open("source.xml")

	if err != nil {
		fmt.Printf("error while opening file %v", err)
		return
	}

	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	c := &Catalog{}

	err = xml.Unmarshal(byteValue, &c)

	if err != nil {
		fmt.Printf("error while unnmarshalling file %v", err)
		return
	}

	catalog := &CatalogJson{Catalog: *c}

	json, err := json.MarshalIndent(catalog, "", "    ")

	if err != nil {
		fmt.Printf("error while MarshalIndent %v", err)
		return
	}

	fmt.Printf("%v", string(json))

	err = os.WriteFile("output.json", json, 0666)

	if err != nil {
		fmt.Printf("error while writing file %v", err)
		return
	}
}
