package main

import (
	"fmt"
	"game/structure"
	"os"
)

var choice string
var ChoicePerso int
func start() {
	fmt.Println("\nM A I N    M E N U")
	fmt.Println("p. Play")
	fmt.Println("r. Return to start screen")
	fmt.Println("q. Quit game")
	
	fmt.Scanln(&choice)
	switch choice {
		case "p":
		submenu1()
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
	fmt.Println("\nC H A R A C T E R S")
	structure.ChoixPersonne()
	fmt.Println("\n r. Return to Main Menu")
	fmt.Println(" q. Quit game")
	
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


func submenu_perso1() {
	fmt.Println("T A N K")
	structure.InfoTank()
	fmt.Println("\ns. Start game as Tank")
	fmt.Println("r. Return to Characters")
	fmt.Println("q. Quit game")

	fmt.Scanln(&choice)
	switch choice {
		case "s":
			ChoicePerso = 1
		submenu_inventaire()
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
	structure.InfoElfe()
	fmt.Println("\ns. Start game as Elfe")
	fmt.Println("r. Return to Characters")
	fmt.Println("q. Quit game")
	
	fmt.Scanln(&choice)
	switch choice {
		case "s":
			ChoicePerso = 2
		submenu_inventaire()
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
	structure.InfoSorcier()
	fmt.Println("\ns. Start game as Sorcier")
	fmt.Println("r. Return to Characters")
	fmt.Println("q. Quit game")
	
	fmt.Scanln(&choice)
	switch choice {
		case "s":
			ChoicePerso = 3
		submenu_inventaire()
		case "r":
		submenu1()
		case "q":
		os.Exit(0)
		default:
		fmt.Println("Incorrect choice.")
		submenu_perso3()
	}
}
func submenu_inventaire() {
	var coins int = 100
    var stock int = 0
    var stockmax int = 5
    fmt.Println("\nI N V E N T O R Y")
    fmt.Println("You have", coins, "gold coins and", stock, "item on", stockmax)
	fmt.Print("\n")
	structure.Obj()
	fmt.Println("\ns. Start")
	fmt.Println("r. Return to Character")
	fmt.Println("q. Quit game")

	var choice string
	fmt.Scanln(&choice)
	switch choice {
		case "s":
			structure.Game()
		case "r":
		submenu_perso1()
		case "q":
		os.Exit(0)
		default:
		fmt.Println("Incorrect choice.")
		submenu_inventaire()
	}
}