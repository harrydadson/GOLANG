package main

import "fmt"

func main() {

	var a [3]int
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(len(a))

	a[0] = 1
	fmt.Println(a)

	b := [4]int{1,2,3,4}
	fmt.Println(b)

	c := [...]int{1,2}
	fmt.Println(c)

	for _, v := range b {
		fmt.Println(v)
	}
	fmt.Println()

	array := [...]int{0: 5}
	fmt.Println(array)

	arr := [...]int{1: 1, 4: 5}
	fmt.Println(arr)

	strArr := [...]string{"string1", "string2"}
	fmt.Println(strArr)

	arr2d := [2][2]int{{1, 2}, {3,4}}
	fmt.Println(arr2d)

	arr3d := [2][2][2]int{arr2d, arr2d}
	fmt.Println(arr3d)
}