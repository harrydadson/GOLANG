package main

import (
	"encoding/json"
	"fmt"
)

type Ninja struct {
	Name string `json: "full_name"`
	Weapon string `json: "weapon"`
	Level int `json: "level"`
}

func main() {

	fmt.Println("Hello, world!")

	sFrom := `{"full_name": "Wallace", "weapon": "Ninja Star", "level": 1}` // json string
	fmt.Println(sFrom)

	var wallace Ninja
	err := json.Unmarshal([]byte(sFrom), &wallace)
	if err != nil {
		fmt.Println("Sad boi")
	}
	fmt.Println(wallace)

	sTo, err := json.Marshal(wallace)
	if err != nil {
		fmt.Println("Sad boi")
	}
	fmt.Println("%T\n", sTo)
	fmt.Println("%s\n", sTo)

}