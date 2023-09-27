package main

import ("fmt"
    "game/structure" //permet l'accès aux fichiers du dossier structure
)

//première fonction lancée
func main() {
    for {
        fmt.Println("\nBienvenue")
        fmt.Println("Click 's' to start or 'q' to quit")  //premier choix du jeu, commencer ou quitter
        var choice string //choice va servir pour les choix
        fmt.Scanln(&choice)
        switch choice {
        	case "s":
            fmt.Println("Goodluck and enjoy !")
            structure.Start()  //lancement du jeu
       		case "q":
            fmt.Println("Goodbye!")
            return //quitter
       		default:
            fmt.Println("Error, please try again")  //choix incorrect, la fonction recommence
        }
    }
}