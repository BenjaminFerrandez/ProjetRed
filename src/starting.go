package main

import (
	"fmt"
	"game/character"
)


func start() {
	fmt.Println("M A I N    M E N U")
	fmt.Println("a. Info perso")
	fmt.Println("b. Inventory")
	fmt.Println("q. Quitter")

	var choice string
	fmt.Print("Enter letter of choice: ")
	fmt.Scanln(&choice)

	switch choice {
		case "a":
		submenu1()
		case "b":
		submenu2()
		case "q":
		return
		default:
		fmt.Println("Incorrect choice. Please select an option from the menu (a, b or q)")
	}
}
func submenu1() {
	fmt.Println("C H A R A C T E R S")
	character.Personne()
	fmt.Print("Back to Main Menu (Press q): ")

	var subchoice1 string
	fmt.Scanln(&subchoice1)

	switch subchoice1 {
		case "q":
		fmt.Println("Return to Main Menu....")
		start()
		default:
		fmt.Println("Incorrect choice.")
		submenu1()
	}
}
func submenu2() {
	fmt.Println("I N V E N T O R Y")
	fmt.Println("You havo 0 items in inventory")
	fmt.Print("Back to Main Menu (Press q): ")

	var subchoice2 string
	fmt.Scanln(&subchoice2)

	switch subchoice2 {
		case "q":
		fmt.Println("Return to Main Menu....")
		start()
		default:
		fmt.Println("Incorrect choice.")
		submenu2()
	}
}