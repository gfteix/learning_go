package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Our function doesn't need to care where or how the printing happens, so we should accept an interface rather than a concrete type.
// fmt.Fprintf is like fmt.Printf but instead takes a Writer to send the string to, whereas fmt.Printf defaults to stdout.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// http.ResponseWriter also implements io.Writer so this is why we could re-use our Greet function inside our handler.

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
