package main

import (
	"fmt"
	"os"
	"text/template"
)

type secret struct {
	Name string
	Secret string
}

func main() {
	fmt.Println("Hello")

	secret := secret{"Wallace", "I'm a Ninja"} // initialized in struct

	templatePath := "/Users/harry.dadson@ibm.com/GoLang/Practice_go/14-Text-template/template.txt"

	t, err := template.New("template.txt").ParseFiles(templatePath)

	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(os.Stdout, secret) // print out to console using Stdout
	if err != nil {
		fmt.Println(err)
	}
}