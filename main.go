package main

import (
	"fmt"
	"strings"
)

type Sail struct {
	brand       string
	name        string
	cost        float32
	size_m2     float32
	year        int
	num_battens int
	has_sailbag bool
	new         bool
	notes       string
}

func AddSail() Sail {
	var new_sail Sail

	fmt.Println("Brand?")
	fmt.Scanln(&new_sail.brand)

	fmt.Println("Name?")
	fmt.Scanln(&new_sail.name)

	fmt.Println("Cost?")
	fmt.Scanln(&new_sail.cost)

	fmt.Println("Size in meters squared?")
	fmt.Scanln(&new_sail.size_m2)

	fmt.Println("Year?")
	fmt.Scanln(&new_sail.year)

	fmt.Println("Number of battens?")
	fmt.Scanln(&new_sail.num_battens)

	fmt.Println("Does it have a sailbag?")
	var bag_input string
	fmt.Scanln(&bag_input)
	bag_input_lower := strings.ToLower(bag_input)
	if bag_input_lower == "yes" {
		new_sail.has_sailbag = true
	}

	fmt.Println("Is it new?")
	var new_conditon string
	fmt.Scanln(&new_conditon)
	new_conditon_lower := strings.ToLower(new_conditon)
	if new_conditon_lower == "yes" {
		new_sail.new = true
	}

	fmt.Println("Do you want to add any notes?")
	fmt.Scanln(&new_sail.notes)

	return new_sail
}

func (s Sail) showSail() {
	fmt.Println("Let's review the new sail\n")
	fmt.Printf("%d %s by %s\n", s.year, s.name, s.brand)
	fmt.Printf("Cost: %.2f\n", s.cost)
	fmt.Printf("Size: %.1f sq meter\n", s.size_m2)
	fmt.Printf("Num battens: %d\n", s.num_battens)
	fmt.Printf("Sailbag: %t\n", s.has_sailbag)
	fmt.Printf("New: %t\n", s.new)
	fmt.Println(s.notes)
}

type Board struct {
	brand      string
	name       string
	cost       float32
	length_cm  int
	liters     int
	inflatable bool
	new        bool
	year       int
	notes      string
}

func AddBoard() Board {
	var new_board Board

	fmt.Println("Brand?")
	fmt.Scanln(&new_board.brand)

	fmt.Println("Name?")
	fmt.Scanln(&new_board.name)

	fmt.Println("Cost?")
	fmt.Scanln(&new_board.cost)

	fmt.Println("Length in centimeters?")
	fmt.Scanln(&new_board.length_cm)

	fmt.Println("Volume in liters?")
	fmt.Scanln(&new_board.liters)

	fmt.Println("Year?")
	fmt.Scanln(&new_board.year)

	fmt.Println("Is it inflatable?")
	var inflatable_input string
	fmt.Scanln(&inflatable_input)
	inflatable_input_lower := strings.ToLower(inflatable_input)
	if inflatable_input_lower == "yes" {
		new_board.inflatable = true
	}

	fmt.Println("Is it new?")
	var new_conditon string
	fmt.Scanln(&new_conditon)
	new_conditon_lower := strings.ToLower(new_conditon)
	if new_conditon_lower == "yes" {
		new_board.new = true
	}

	fmt.Println("Do you want to add any notes?")
	fmt.Scanln(&new_board.notes)

	return new_board
}

func (b Board) showBoard() {
	fmt.Println("Let's review the new board\n")
	fmt.Printf("%d %s by %s\n", b.year, b.name, b.brand)
	fmt.Printf("Cost: %.2f\n", b.cost)
	fmt.Printf("Length: %d cm.  Liters: %d\n", b.length_cm, b.liters)
	fmt.Printf("Inflatables: %t\n", b.inflatable)
	fmt.Printf("Condition: %t\n", b.new)
	fmt.Println(b.notes)
}

func main() {
	add_another := true
	for add_another == true {
		fmt.Println("Do you want to add something? Sail or a board?")
		var selection string
		fmt.Scanln(&selection)
		selection_lower := strings.ToLower(selection)

		if selection_lower == "sail" {
			fmt.Println("Ok let's add a sail")
			a_sail := AddSail()
			a_sail.showSail()
		} else if selection_lower == "board" {
			fmt.Println("Ok let's add a board")
			a_board := AddBoard()
			a_board.showBoard()
		} else if selection_lower == "no" {
			add_another = false
			fmt.Println("Ok bye")
		} else {
			fmt.Println("Unrecognized input. Select sail or board.")
		}
	}
}
