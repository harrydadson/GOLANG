package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
/*	
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Printf("%+v", alex)
*/
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contact: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// &jim - give me the memory address of the value this variable point at

	jim.updateName("Jimmy")
	jim.print()
}

	// *pointer - give mne the value this memory address is pointing at
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

/*
001			person{firstName: "Jim".....}
address			value

Turn address into value with *address
Turn value into address with &value
*/

/*
int, float, string, bool, structs are Value Types therefore use pointers to
change these in a function

slices, maps, channels, pointers, and functions are Reference Types therefore
we dont worry about using pointers
*/
