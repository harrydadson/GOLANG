package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the year you were born: ")
	sc.Scan()
	input, _ := strconv.ParseInt(sc.Text(), 10, 64)
	fmt.Printf("You will be %d years \n", 2021 - input)
}