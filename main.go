package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	maxHealth = 3
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var name string
	var character string
	health := 3
	var characters [3]string = [3]string{
		"Halfling (+1 vs. rats & dragons)",
		"Dwarf (+1 vs. goblins & trolls)",
		"Elf (+1 vs. orcs & skeletons)",
	}
	var items []string

	fmt.Println("Hello, World!")
	fmt.Println("What is your name?")
	fmt.Scanln(&name)

	fmt.Println("Let's roll your character. Press enter.")
	fmt.Scanln()
	characterRoll := d6()
	switch characterRoll {
	case 1, 2:
		character = characters[0]
	case 3, 4:
		character = characters[1]
	case 5, 6:
		character = characters[2]
	}
	fmt.Println(characterRoll, "!", "You rolled:", character)
	fmt.Println("Hello", name+"!", "Type 'c' to see you character sheet")
	var command string
	fmt.Scan(&command)
	if command == "c" {
		printStatus(name, character, health, items)
	}
}

func printStatus(name string, character string, health int, items []string) {
	fmt.Println(name+",", character)
	fmt.Println("Your current health:", health, "/", maxHealth)
	fmt.Print("Items in your inventory: ")
	fmt.Printf("%v", items)
	fmt.Println()
}

func d6() int {
	return rand.Intn(6-1) + 1
}
