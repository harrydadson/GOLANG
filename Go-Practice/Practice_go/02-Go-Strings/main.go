package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	s := "Hello world"
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(s[0])
	fmt.Printf("%c", s[0])
	fmt.Println(s[0:6])
	fmt.Println(s[6:])

	s = s + " Again"
	s += " Again"
	fmt.Println(s)
	fmt.Println("Hello \nWorld")
	fmt.Println("Hello \tWorld")
	fmt.Println("Hello \bWorld")

	abc := "abc"
	b := []byte(abc)
	fmt.Printf("%s, %s", abc, b)

	// String Lib
	fmt.Println(strings.ToUpper("Hello Harry"))
	fmt.Println(strings.ToLower("Hello Harry"))
	fmt.Println(strings.HasPrefix("Hello Harry", "Hello")) // True
	fmt.Println(strings.HasSuffix("Hello Harry", "Hello")) // false
	fmt.Println(strings.HasSuffix("Hello Harry", "Harry")) // false
	fmt.Println(strings.Replace("Hello Harry", "Harry", "Amy", 1))
	fmt.Println(strings.Replace("Hello Harry Harry", "Harry", "Manny", 2))

	// String Builder - tool to build your strings
	var sb strings.Builder
	fmt.Println("This is a String Builder: ", sb.String())

	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())

	sb.WriteString("Hello")
	fmt.Println("This is a String Builder: ", sb.String())

	fmt.Println(sb.Cap()) // 0 2 4 8 <- space
	fmt.Println(sb.Len()) // just length of string

	sb.Write([]byte{65, 65, 65})
	fmt.Println(sb.String())

	x := 123
	y := strconv.Itoa(x)
	fmt.Printf("%T", y)
	fmt.Println()

	z,_ := strconv.Atoi(y)
	fmt.Println(z)
	fmt.Printf("%T", z)
}
