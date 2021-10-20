package main

import (
	"errors"
	"fmt"
)

func main() {
	Hello()
	variadic(1,2,3)
	fmt.Println(sum(1, 2))
	fmt.Println(nameLength("Wallace"))

	result, error := greeting(" ")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(result)
	}

	resultOk, ok := greetingOK("")
	if ok {
		fmt.Println(resultOk)
	} 

	fmt.Println(fibonnacci(10))

	i := 1
	fmt.Println(i)
	//withoutPointer(i)
	fmt.Println(i)
	withoutPointer(&i)
	fmt.Println(i)
}

func Hello() {
	fmt.Println("Hello World")
}

func variadic(integers ...int) {
	for _,i := range integers {
		fmt.Println(i)
	}
}

func sum(a, b int) int {
	return a + b
}

func nameLength(name string) (string, int) {
	return name, len(name)
}

func greeting(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("Empty name")
	} else {
		return "Hello " + name, nil
	}
}

func greetingOK(name string) (string, bool) {
	if len(name) == 0 {
		return "", false
	} else {
		return "Hello " + name, true
	}
}
	// Recursive func
func fibonnacci(i int) int {
	if i == 0 {
		return 0
	} else if i == 1 {
		return 1
	} else {
		return fibonnacci(i -1) + fibonnacci(i - 2)
	}
}

func withoutPointer(i *int) {
	*i = *i + 1
}
