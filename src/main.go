package main

import (
	"fmt"
	"game/structure" //permet l'accès aux fichiers du dossier structure

	"github.com/hajimehoshi/ebiten" //import de bibliotheque ebiten pour graphique
	"golang.org/x/image/font"
)

const ( //taille de l'écran
	screenWidth  = 1024
	screenHeight = 768
)

var ( //caractéristique de l'affichage
	screen          *ebiten.Image
	backgroundImg   *ebiten.Image
	gameStarted     bool
	startScreenText string
	Font            font.Face
)

//première fonction lancée
func main() {
	for {
		ebiten.Run(history, screenWidth, screenHeight, 1, "Death Back") //titre

		fmt.Println("\nClick 's' to start or 'q' to quit") //premier choix du jeu, commencer ou quitter
		var choice string                                  //choice va servir pour les choix
		fmt.Scanln(&choice)
		switch choice {
		case "s":
			fmt.Println("Goodluck and enjoy !")
			structure.Start() //lancement du jeu
		case "q":
			fmt.Println("Goodbye!")
			return //quitter
		default:
			fmt.Println("Error, please try again") //choix incorrect, la fonction recommence
		}
	}
}
