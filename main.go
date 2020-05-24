package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sail_fileloc string = "./data/saved_sails.csv"
var board_fileloc string = "./data/saved_boards.csv"
var sail_header = []string{"Brand", "Name", "Year", "Cost", "Size (Sq Meter)", "Num Battens", "Has Sailbag", "New", "Notes"}
var board_header = []string{"Brand", "Name", "Year", "Cost", "Length (cm)", "Liters", "Inflatable", "New", "Notes"}

func CreateFileIfNotExists(sail_or_board string, fileloc string) bool {
	success := true

	if _, err := os.Stat(fileloc); os.IsNotExist(err) {
		file, err := os.Create(fileloc)
		if err != nil {
			fmt.Printf("Could not create file at %s", fileloc)
			fmt.Println(err)
			success = false
		}
		defer file.Close()

		if success {
			writer := csv.NewWriter(file)
			defer writer.Flush()

			if sail_or_board == "sail" {
				err := writer.Write(sail_header)
				if err != nil {
					fmt.Printf("Could not write header to sail file at %s", fileloc)
					fmt.Println(err)
					success = false
				}
			} else {
				err := writer.Write(board_header)
				if err != nil {
					fmt.Printf("Could not write header to board file at %s", fileloc)
					fmt.Println(err)
					success = false
				}
			}
		}
	}
	return success
}

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
	fmt.Printf("Cost: $%.2f\n", s.cost)
	fmt.Printf("Size: %.1f sq meter\n", s.size_m2)
	fmt.Printf("Num battens: %d\n", s.num_battens)
	fmt.Printf("Sailbag: %t\n", s.has_sailbag)
	fmt.Printf("New: %t\n", s.new)
	fmt.Println(s.notes)
}

func saveSailToFile(s Sail) {
	newline := []string{
		s.brand, s.name, strconv.Itoa(s.year), strconv.FormatFloat(float64(s.cost), 'f', 2, 32),
		strconv.FormatFloat(float64(s.size_m2), 'f', 2, 32), strconv.Itoa(s.num_battens),
		strconv.FormatBool(s.has_sailbag), strconv.FormatBool(s.new), s.notes}

	file, err := os.OpenFile(sail_fileloc, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()

	if err != nil {
		fmt.Println("Could not open sail file in append mode")
		fmt.Println(err)
		return
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()
	write_err := writer.Write(newline)

	if write_err != nil {
		fmt.Println("Could not append line to sail file")
		fmt.Println(write_err)
	}
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
	fmt.Printf("Cost: $%.2f\n", b.cost)
	fmt.Printf("Length: %d cm.  Liters: %d\n", b.length_cm, b.liters)
	fmt.Printf("Inflatable: %t\n", b.inflatable)
	fmt.Printf("New: %t\n", b.new)
	fmt.Println(b.notes)
}

func saveBoardToFile(b Board) {
	newline := []string{
		b.brand, b.name, strconv.Itoa(b.year), strconv.FormatFloat(float64(b.cost), 'f', 2, 32),
		strconv.Itoa(b.length_cm), strconv.Itoa(b.liters),
		strconv.FormatBool(b.inflatable), strconv.FormatBool(b.new), b.notes}

	file, err := os.OpenFile(board_fileloc, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()

	if err != nil {
		fmt.Println("Could not open board file in append mode")
		fmt.Println(err)
		return
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()
	write_err := writer.Write(newline)

	if write_err != nil {
		fmt.Println("Could not append line to board file")
		fmt.Println(write_err)
	}
}

func main() {
	sail_file_success := CreateFileIfNotExists("sail", sail_fileloc)
	board_file_success := CreateFileIfNotExists("board", board_fileloc)

	if !sail_file_success || !board_file_success {
		os.Exit(1)
	}

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
			saveSailToFile(a_sail)
		} else if selection_lower == "board" {
			fmt.Println("Ok let's add a board")
			a_board := AddBoard()
			a_board.showBoard()
			saveBoardToFile(a_board)
		} else if selection_lower == "no" {
			add_another = false
			fmt.Println("Ok bye")
		} else {
			fmt.Println("Unrecognized input. Select sail or board.")
		}
	}
}
