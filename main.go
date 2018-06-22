package main

import (
	"fmt"
	"os"
	"net/http"
)

func main() {
	r, err := http.Get("https://bseed.42.fr/run/")
	if err != nil {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	fmt.Println(r)
}
