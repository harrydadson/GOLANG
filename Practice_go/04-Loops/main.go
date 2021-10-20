package main

import "fmt"

func main() {

	for i := 0; i < 5; i += 1{
		fmt.Println(i)
	}

	for i, j := 0,0; i < 5 && j < 5; i, j = i + 1, j + 1 {
		fmt.Println(i, j)
	}

	// For each loops
	s := "Hello World"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d %c", i, s[i])
		fmt.Println()
	} 

	for i, c := range s {
		fmt.Printf("%d %c", i, c)
		fmt.Println()
	}

	// Break
	for _, str := range s {
		if str == ' ' {
			break // vs continue
		}
		fmt.Printf("%c", str)
	}
	fmt.Println()

	// label
	iForLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break iForLoop
			}
			fmt.Printf("%d%d", i, j)
		}
		fmt.Println()
	}
}