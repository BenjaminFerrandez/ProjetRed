package main

import (
	"fmt"
	"game/structure"
	"os"
)

func start() {
	fmt.Println("M A I N    M E N U")
	fmt.Println("a. Info perso")
	fmt.Println("b. Inventory")
	fmt.Println("r. Return to start screen")
	fmt.Println("q. Quit game")
	
	var choice string
	fmt.Scanln(&choice)
	switch choice {
		case "a":
		submenu1()
		case "b":
		submenu2()
		case "r":
		return
		case "q":
		os.Exit(0)
		default:
		fmt.Println("Incorrect choice. Please select an option from the menu (a, b r or q)")
		start()
	}
}

func submenu1() {
	fmt.Println("C H A R A C T E R S")
	structure.Personne()
	fmt.Println("r. Return to Main Menu")
	fmt.Println("q. Quit game")
	
	var subchoice1 string
	fmt.Scanln(&subchoice1)
	switch subchoice1 {
		case "r":
		start()
		case "q":
		os.Exit(0)
		default:
		fmt.Println("Incorrect choice.")
		submenu1()
	}
}

func submenu2() {
	fmt.Println("I N V E N T O R Y")
	fmt.Println("You havo 0 items in inventory")
	fmt.Println("r. Return to Main Menu")
	fmt.Println("q. Quit game")

	var subchoice2 string
	fmt.Scanln(&subchoice2)
	switch subchoice2 {
		case "r":
		start()
		case "q":
		os.Exit(0)
		default:
		fmt.Println("Incorrect choice.")
		submenu2()
	}
}