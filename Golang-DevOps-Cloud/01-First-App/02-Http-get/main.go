package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage: ./02-Http-get ,url.\n")
		os.Exit(1)
	}

	if _,err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("URL is invalid dormat: %v", err)
		os.Exit(1)
	}

	response, err := http.Get(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("HTTP Status Code: %d\nBody: %s\n", response.StatusCode, string(body))

}	