package main

import (
	"fmt"
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
	if bag_input == "yes" {
		new_sail.has_sailbag = true
	}

	fmt.Println("Is it new?")
	var new_conditon string
	fmt.Scanln(&new_conditon)
	if new_conditon == "yes" {
		new_sail.new = true
	}

	fmt.Println("Do you want to add any notes?")
	fmt.Scanln(&new_sail.notes)

	return new_sail
}

func (s Sail) showSail() {
	fmt.Printf("%d %s by %s\n", s.year, s.name, s.brand)
	fmt.Printf("Cost: %.2f\n", s.cost)
	fmt.Printf("Size: %.1f sq meter\n", s.size_m2)
	fmt.Printf("Num battens: %d\n", s.num_battens)
	fmt.Printf("Sailbag: %t\n", s.has_sailbag)
	fmt.Printf("New: %s\n", s.new)
	fmt.Println(s.notes)
}

type Board struct {
	brand      string
	name       string
	cost       float32
	length_cm  int
	liters     int
	inflatable bool
	condition  string
	year       int
	notes      string
}

func (b Board) showBoard() {
	fmt.Printf("%d %s by %s\n", b.year, b.name, b.brand)
	fmt.Printf("Cost: %.2f\n", b.cost)
	fmt.Printf("Length: %d cm.  Liters: %d\n", b.length_cm, b.liters)
	fmt.Printf("Inflatables: %t\n", b.inflatable)
	fmt.Printf("Condition: %s\n", b.condition)
	fmt.Println(b.notes)
}

type Kit struct {
	Sail
	Board
}

func main() {
	fmt.Println("Do you want to add a sail or a board?")
	var selection string
	fmt.Scanln(&selection)
	if selection == "sail" {
		fmt.Println("Ok let's add a sail")
		a_sail := AddSail()
		a_sail.showSail()
	} else if selection == "board" {
		fmt.Println("Ok let's add a board")
	} else {
		fmt.Println("Unrecognized input. Select sail or board.")
	}

}
