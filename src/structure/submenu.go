package structure

import (
	"fmt"

	"os"
)

type Enemy struct {
	Name   string
	Health int
}

//choix du personnage
func submenu1() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("\nC H A R A C T E R S")
	ChoixPersonne() //affiche les noms et classe de chaque personnages
	fmt.Println("\n r. Return to Main Menu")
	fmt.Println(" q. Quit game")

	fmt.Scanln(&Choice)
	switch Choice {
	case "1":
		submenu_perso1() //perso1
	case "2":
		submenu_perso2() //perso2
	case "3":
		submenu_perso3() //perso3
	case "r":
		Start() //retour en arrière
	case "q":
		os.Exit(0) //Quitter
	default:
		fmt.Println("Incorrect Choice.")
		submenu1() //choix incorrect
	}
}

//perso1
func submenu_perso1() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("T A N K")
	InfoTank() //affiche les infos du perso1
	fmt.Println("\ns. Start game as Tank")
	fmt.Println("r. Return to Characters")
	fmt.Println("q. Quit game")

	fmt.Scanln(&Choice)
	switch Choice {
	case "s":
		ChoicePerso = 1
		Obj() //commencer le jeu avec le perso1
	case "r":
		submenu1() //retour en arrière
	case "q":
		os.Exit(0) //Quitter
	default:
		fmt.Println("Incorrect Choice.")
		submenu_perso1() //choix incorrect
	}
}

//perso2
func submenu_perso2() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("E L F E")
	InfoElfe() //affiche les infos du perso2
	fmt.Println("\ns. Start game as Elfe")
	fmt.Println("r. Return to Characters")
	fmt.Println("q. Quit game")

	fmt.Scanln(&Choice)
	switch Choice {
	case "s":
		ChoicePerso = 2
		Obj() //commencer le jeu avec le perso2
	case "r":
		submenu1() //retour en arrière
	case "q":
		os.Exit(0) //Quitter
	default:
		fmt.Println("Incorrect Choice.")
		submenu_perso2() //choix incorrect
	}
}

//perso3
func submenu_perso3() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("S O R C I E R")
	InfoSorcier() //affiche les infos du perso3
	fmt.Println("\ns. Start game as Sorcier")
	fmt.Println("r. Return to Characters")
	fmt.Println("q. Quit game")

	fmt.Scanln(&Choice)
	switch Choice {
	case "s":
		ChoicePerso = 3
		Obj() //commencer le jeu avec le perso3
	case "r":
		submenu1() //retour en arrière
	case "q":
		os.Exit(0) //Quitter
	default:
		fmt.Println("Incorrect Choice.")
		submenu_perso3() //choix incorrect
	}
}
