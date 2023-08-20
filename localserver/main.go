package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: localserver <directory>")
		return
	}

	dir := os.Args[1]
	fs := http.FileServer(http.Dir(dir))

	http.Handle("/", fs)
	fmt.Println("Serving on :8080...")
	http.ListenAndServe(":8080", nil)
}
