package main

import ("fmt"
    "game/structure" //permet l'accès aux fichiers du dossier structure

    "github.com/hajimehoshi/ebiten"
)
const (
    screenWidth  = 800
    screenHeight = 600
)
var screen *ebiten.Image
//première fonction lancée
func main() {
    for {
        ebiten.Run(history, screenWidth, screenHeight, 1, "Death Back")
      
        fmt.Println("D E A T H   B A C K") //titre
 
 

        fmt.Println("\nClick 's' to start or 'q' to quit") //premier choix du jeu, commencer ou quitter
        var choice string //choice va servir pour les choix
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
