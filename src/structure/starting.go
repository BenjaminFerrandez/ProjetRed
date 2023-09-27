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
	fmt.Println("p. Play")  //différents choix = jouer, retour en arrière ou quitter
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
		fmt.Println("Incorrect Choice. Please select an option from the menu (a, b r or q)")
		Start() //choix incorrect, la fonction recommence
	}
}

//choix du personnage
func submenu1() {
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
	fmt.Println("T A N K")
	InfoTank() //affiche les infos du perso1
	fmt.Println("\ns. Start game as Tank")
	fmt.Println("r. Return to Characters")
	fmt.Println("q. Quit game")

	fmt.Scanln(&Choice)
	switch Choice {
		case "s":
			ChoicePerso = 1
		submenu_inventaire() //commencer le jeu avec le perso1
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
	fmt.Println("E L F E")
	InfoElfe() //affiche les infos du perso2
	fmt.Println("\ns. Start game as Elfe")
	fmt.Println("r. Return to Characters")
	fmt.Println("q. Quit game")
	
	fmt.Scanln(&Choice)
	switch Choice {
		case "s":
			ChoicePerso = 2
		submenu_inventaire()  //commencer le jeu avec le perso2
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
	fmt.Println("S O R C I E R")
	InfoSorcier() //affiche les infos du perso3
	fmt.Println("\ns. Start game as Sorcier")
	fmt.Println("r. Return to Characters")
	fmt.Println("q. Quit game")
	
	fmt.Scanln(&Choice)
	switch Choice {
		case "s":
			ChoicePerso = 3
		submenu_inventaire()  //commencer le jeu avec le perso3
		case "r":
		submenu1() //retour en arrière
		case "q":
		os.Exit(0) //Quitter
		default:
		fmt.Println("Incorrect Choice.") //choix incorrect
		submenu_perso3() 
	}
}

//initialisation de l'inventaire
func submenu_inventaire() {
	var coins int = 100 //on démarre avec 100 pièces,
    var stock int = 0 // 0 objet
    var stockmax int = 5 // et on peut avoir 5 objet au maximum
    fmt.Println("\nI N V E N T O R Y")
    fmt.Println("You have", coins, "gold coins and", stock, "item on", stockmax)
	fmt.Print("\n")
	Obj() // affiche la liste des objets du jeu
}