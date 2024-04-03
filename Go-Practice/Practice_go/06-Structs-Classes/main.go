package main

import "fmt"

func main() {

	type ninja struct {
		name string
		weapons []string
		level int
	}

	//wallace := ninja{name: "Wallace"}
	wallace := ninja{"Wallace", []string{"Ninja Star"}, 1}
	fmt.Println(wallace)
	fmt.Println(wallace.name)
	fmt.Println(wallace.weapons)
	fmt.Println(wallace.level)
	wallace.level++
	wallace.weapons = append(wallace.weapons, "Ninja Sword")

	type dojo struct {
		name string
		ninja ninja
	}

	goLangDojo := dojo {
		name: "Golang Dojo",
		ninja: wallace,
	}
	fmt.Println(goLangDojo)
	fmt.Println(goLangDojo.ninja.level)

	goLangDojo.ninja.level = 3
	fmt.Println(goLangDojo.ninja.level)
	fmt.Println(wallace.level)

	type addressedDojo struct {
		name string
		ninja *ninja
	}
	addressed := addressedDojo{"Addressed goland dojo", &wallace}
	fmt.Println(addressed)
	fmt.Println(*addressed.ninja)

	addressed.ninja.level = 4
	fmt.Println(addressed.ninja.level)
	fmt.Println(wallace.level)

	johnnyPointer := new(ninja)
	fmt.Println(johnnyPointer)
	fmt.Println(*johnnyPointer)
	johnnyPointer.name = "Johnny"
	johnnyPointer.weapons = []string{"Ninja Star"}
	johnnyPointer.level = 1
	fmt.Println(*johnnyPointer)

	intern := ninjaIntern{name: "Intern", level: 1}
	intern.sayHello()
	intern.sayName()

}

type ninjaIntern struct {
	name string
	level int
}

func (ninjaIntern) sayHello() {
	fmt.Println("Hello")
}

func (n ninjaIntern) sayName() {
	fmt.Println("Hello", n.name)
}





