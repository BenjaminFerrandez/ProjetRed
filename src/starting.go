package main

import (
	"fmt"
	"game/structure"
	"os"
)

var choice string

func start() {
	fmt.Println("M A I N    M E N U")
	fmt.Println("a. Info perso")
	fmt.Println("b. Inventory")
	fmt.Println("r. Return to start screen")
	fmt.Println("q. Quit game")
	
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
	
	fmt.Scanln(&choice)
	switch choice {
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
	
	fmt.Scanln(&choice)
	switch choice {
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
	structure.InitTank()
	fmt.Println("r. Return to Characters")
	fmt.Println("q. Quit game")

	fmt.Scanln(&choice)
	switch choice {
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
	structure.InitElfe()
	fmt.Println("r. Return to Characters")
	fmt.Println("q. Quit game")
	
	fmt.Scanln(&choice)
	switch choice {
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
	structure.InitSorcier()
	fmt.Println("r. Return to Characters")
	fmt.Println("q. Quit game")
	
	fmt.Scanln(&choice)
	switch choice {
		case "r":
		submenu1()
		case "q":
		os.Exit(0)
		default:
		fmt.Println("Incorrect choice.")
		submenu_perso3()
	}
}