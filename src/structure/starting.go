package structure

import (
	"fmt"

	"os"
)

var Choice string   //va servir pour les choix de menu
var ChoicePerso int //va servir pour choisir son personnage

//démarrage
func Start() {
	fmt.Println("\nM A I N    M E N U")
	//différents choix = jouer, retour en arrière ou quitter
	fmt.Println("\np. Play")
	fmt.Println("q. Quit game")

	fmt.Scanln(&Choice)
	switch Choice {
	case "p":
		submenu1() //suite du jeu
	case "q":
		os.Exit(0) //Quitter
	default:
		fmt.Println("Incorrect Choice. Please select an option from the menu (p, r or q)")
		Start() //choix incorrect, la fonction recommence
	}
}
