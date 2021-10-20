package main

import "fmt"

func main() {
	i, j := 42, 2701
	fmt.Println(i, j)
	
	p := &i
	fmt.Println(*p) // Dereferencing
	fmt.Printf("%T\n", p)

	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37

	wallace := "Uses Ninja star"
	fmt.Println(wallace)

	changeGear(wallace)
	fmt.Println(wallace)

	changeGearPointer(&wallace)
	fmt.Println(wallace)

	// & is get the pointer, * is get the reference or dereference
	x := 7
	fmt.Println(&x) // where is location var x sits, so 0x123

	y := &x // y is location of x 0x123
	fmt.Println(x, y) // 7 0x123

	*y = 8 // change what x is pointing to
	fmt.Println(x, y) // 8 0x123

	toChange := "Hello"
	var pntr *string = &toChange
	fmt.Println(*pntr) // Hello
	fmt.Println(&pntr) // 0x123
}

func changeGear(wallaceCopy string) {
	wallaceCopy = "Uses Ninja sword"
}

func changeGearPointer(wallacePointer *string) {
	*wallacePointer = "Uses Ninja Pointer"
}