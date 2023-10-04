package structure

import (
	"fmt"
	
	"os"
)

var Choice string
var ChoicePerso int //va servir pour choisir son personnage

//démarrage
func Start() {
	fmt.Println("\nM A I N    M E N U")
	fmt.Println("\np. Play") //différents choix = jouer, retour en arrière ou quitter
	fmt.Println("r. Return to start screen")
	fmt.Println("q. Quit game")

	fmt.Scanln(&Choice)
	switch Choice {
		case "p":
			submenu1() //suite du jeu
		case "r":
			return //retour en arrière
		case "q":
			os.Exit(0) //Quitter
		default:
			fmt.Println("Incorrect Choice. Please select an option from the menu (p, r or q)")
			Start() //choix incorrect, la fonction recommence
	}
}