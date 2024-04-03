package main

import (
	"fmt"
	"math"
)

func main() {

	const five int = 5
	fmt.Println(five)

	const pi = 3.14
	fmt.Println(math.Pi)

	const (
		a = 1
		b 
		c = 3
		d 
	)
	fmt.Println(a, b, c, d)

	const (
		zero int = iota
		one
		two
	)
	fmt.Println(zero, one, two)

	const (
		_ = 1 << (10 * iota) // bitwise shift to the left
		kb
		mb
		gb
		tb
		pb
		eb
		zb
		yb
	)
	fmt.Println(kb, mb, gb)
}