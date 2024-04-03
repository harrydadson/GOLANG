package main

import "fmt"

func main() {
	
	var mp map[string]int = map[string]int{ // or mp :=

		"apple": 5,
		"pear": 6,
		"orange": 9,
	}
	fmt.Println(mp["apple"])
	mp["apple"] = 900
	fmt.Println(mp)

	mp["tim"] = 24

	delete(mp, "orange")
	fmt.Println(mp)

	// check if key exists
	val, ok := mp["apple"]
	fmt.Println(val, ok)
	fmt.Println(mp)

	fmt.Println(len(mp))




}