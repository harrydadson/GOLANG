package main

import "fmt"

func test2(myFunc func(int) int) {
	fmt.Println(myFunc(7))
}

// function closure
func returnFunc(x string) func()  {
	return func() { fmt.Println(x) }
}

func main() {

	test := func(x int)  int {
		return x * -1
	}
	// test(5) // or just (5)
	// fmt.Println(test)
	test3 := func(x int) int{
		return x * 7
	} 
	test2(test)
	test2(test3)

	returnFunc("Hello")() // Hello
	x := returnFunc("goodbye")
	x() // goodbye
}
