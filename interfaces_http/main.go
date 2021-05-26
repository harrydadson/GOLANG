package main

import (
	"fmt" 
	"os"
	"net/http"
	"io"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error is:", err)
		os.Exit(1)
	}
/*	
	bs := make([]byte, 99999) // type of string
	resp.Body.Read(bs)
	
	fmt.Println(string(bs))
*/

// Simplified Version

//	io.Copy(os.Stdout, resp.Body) // per say (writerInterface, readerInterface)

	lw := logWriter{} // custom made
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes", len(bs))
	return len(bs), nil
}