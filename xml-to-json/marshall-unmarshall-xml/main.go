package main

import (
	"encoding/xml"
	"fmt"
)

type MyXML struct {
	Cat `xml:"cat"`
}

type Cat struct {
	Name string `xml:"name"`
	Age  uint   `xml:"age"`
}

type Product struct {
	ID       uint64   `xml:"id"`
	Name     string   `xml:"name"`
	SKU      string   `xml:"sku"`
	Price    float64  `xml:"price"`
	Category Category `xml:"category"`
}
type Category struct {
	ID   uint64 `xml:"id"`
	Name string `xml:"name"`
}

func main() {
	myXML := []byte(`<cat>
    <name>Ti</name>
    <age>23</age>
</cat>`)
	c := MyXML{}
	err := xml.Unmarshal(myXML, &c)
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Cat.Name)
	fmt.Println(c.Cat.Age)

	// Encoding XML is also very easy; you can use the xml.Marshal, or like here, the function xml.MarshalIndent to display a pretty xml string
	p := Product{ID: 42, Name: "Tea Pot", SKU: "TP12", Price: 30.5, Category: Category{ID: 2, Name: "Tea"}}
	bI, err := xml.MarshalIndent(p, "", "   ")
	if err != nil {
		panic(err)
	}

	/*
		This header will give important information to the systems that will use your xml document. It gives information about the version of XML used (“1.0” in our case that dates back to 1998) and the encoding of your document.
		The parser will need this information to decode your document correctly.
		The xml package defines a constant that you can use directly: xml.Header
	*/
	xmlWithHeader := xml.Header + string(bI)
	fmt.Println(xmlWithHeader)

	/*
		xml.Name type

		You can add a special field to your struct to control the name of the XML element :

		type MyXmlElement struct {
			XMLName xml.Name `xml:"myOwnName"`
			Name  string `xml:"name"`
		}


		Specificity of XML tags


	*/
}
