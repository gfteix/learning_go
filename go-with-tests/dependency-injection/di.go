package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Our function doesn't need to care where or how the printing happens, so we should accept an interface rather than a concrete type.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

// http.ResponseWriter also implements io.Writer so this is why we could re-use our Greet function inside our handler.
func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
