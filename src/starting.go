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
	structure.ChoixPersonne()
	fmt.Println("r. Return to Main Menu")
	fmt.Println("q. Quit game")
	
	var subchoice1 string
	fmt.Scanln(&subchoice1)
	switch subchoice1 {
		case "1":
		submenu_perso1() 
		case "2":
		submenu_perso2() 
		case "3":
		submenu_perso3() 
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
func submenu_perso1() {
	fmt.Println("T A N K")
	structure.TankInfo()
	fmt.Println("r. Return to Characters")
	fmt.Println("q. Quit game")
	
	var subchoice3 string
	fmt.Scanln(&subchoice3)
	switch subchoice3 {
		case "r":
		submenu1()
		case "q":
		os.Exit(0)
		default:
		fmt.Println("Incorrect choice.")
		submenu_perso1()
	}
}
func submenu_perso2() {
	fmt.Println("E L F E")
	structure.ElfeInfo()
	fmt.Println("r. Return to Characters")
	fmt.Println("q. Quit game")
	
	var subchoice4 string
	fmt.Scanln(&subchoice4)
	switch subchoice4 {
		case "r":
		submenu1()
		case "q":
		os.Exit(0)
		default:
		fmt.Println("Incorrect choice.")
		submenu_perso2()
	}
}
func submenu_perso3() {
	fmt.Println("S O R C I E R")
	structure.SorcierInfo()
	fmt.Println("r. Return to Characters")
	fmt.Println("q. Quit game")
	
	var subchoice5 string
	fmt.Scanln(&subchoice5)
	switch subchoice5 {
		case "r":
		submenu1()
		case "q":
		os.Exit(0)
		default:
		fmt.Println("Incorrect choice.")
		submenu_perso3()
	}
}