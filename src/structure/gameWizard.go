package structure

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

//caractéristique du perso3
type Wizard struct {
	Name   string //son nom
	Health int    //ses points de vie de départ
}

//caractéristique des attaques du perso3
type WizardAttack struct {
	Name   string //le nom
	Damage int    //les dégats
}

//caractéristique de l'inventaire du perso3
type InventoryWizard struct {
	Name       string //nom de l'objet
	EffectType string //son type
	Value      int
}

var (
	enemyWizard  = Wizard{Name: "Enemy", Health: 75}   //nom et points de vie du perso3
	playerWizard = Wizard{Name: "Reicros", Health: 90} //nom et points de vie de l'ennemi
	ValueWizard  int
	DefeatWizard int
	WinWizard    string
)

//début du combat en tant que perso3
func GameWizard(inventory []Object) {
	fmt.Print("\033[H\033[2J")
	rand.Seed(time.Now().UnixNano())

	if DefeatWizard != 0 { //augmente les HP de l'ennemi si on l'a déjà tuer
		playerWizard.Health = 90
		enemyWizard.Health = 75
		DefeatWizard = 0
		ValueWizard = 0
	}

	attack1 := WizardAttack{Name: "Fireball ", Damage: 20}
	attack2 := WizardAttack{Name: "Lightning", Damage: 30}

	fmt.Println("Welcome to game!")

	Round := 1

	for playerWizard.Health > 0 && enemyWizard.Health > 0 { //vérifie si les 2 ont de la vie
		fmt.Printf("\nRound: %d\n", Round)
		fmt.Printf("%s (HP: %d) vs %s (HP: %d)\n", playerWizard.Name, playerWizard.Health, enemyWizard.Name, enemyWizard.Health)
		var choice int
		if len(inventory) == 0 {
			fmt.Println("You have no items left.")
			choice = 1
		} else {
			fmt.Println("Choose action:")
			fmt.Println("1. Attack")
			fmt.Println("2. Use item")
			fmt.Print("Enter your choice: ")
			fmt.Scanln(&choice)
		}

		switch choice {
		case 1: //1 = attaquer
			fmt.Println("Choose attack:")
			fmt.Printf("1. %s (Damage: %d)\n", attack1.Name, attack1.Damage)
			fmt.Printf("2. %s (Damage: %d)\n", attack2.Name, attack2.Damage)

			var attackChoice int
			fmt.Print("Enter number of attack: ")
			fmt.Scanln(&attackChoice)

			var playerAttack WizardAttack
			var enemyDamage int
			switch attackChoice {
			case 1: //1.1 = attaque 1
				playerAttack = attack1
				enemyDamage = playerAttack.Damage
			case 2: //1.2 = attaque 2
				if rand.Float64() < 0.6 {
					playerAttack = attack2
					enemyDamage = attack2.Damage
				} else {
					enemyDamage = rand.Intn(attack2.Damage)
				}
			default:
				fmt.Println("Incorrect choice of attack.")
				continue
			}

			enemyWizard.Health -= enemyDamage
			fmt.Printf("\nYou did %d damage to enemy!\n", enemyDamage)

		case 2: //2 = utiliser un objet
			fmt.Println("Choose item to use:")
			for i, item := range inventory {
				fmt.Printf("%d. %s\n", i+1, item.Name)
			}

			var itemChoice int
			fmt.Print("Enter number of item: ")
			fmt.Scanln(&itemChoice)

			if itemChoice < 1 || itemChoice > len(inventory) {
				fmt.Println("Incorrect choice of item.")
				continue
			}

			usedItem := inventory[itemChoice-1]
			switch usedItem.Effect {
			case "Heal":
				playerWizard.Health += 15
				fmt.Printf("\nYou used %s and restored %d health!\n", usedItem.Name, 15)
			case "Poison":
				enemyWizard.Health -= 15
				fmt.Printf("\nYou used %s and dealt %d damage to the enemy!\n", usedItem.Name, 15)
			case "Upgrade":
				playerAttack := attack1
				playerAttack.Damage = int(float64(playerAttack.Damage) * 1.2)
				enemyDamage := rand.Intn(playerAttack.Damage)
				enemyWizard.Health -= enemyDamage
				fmt.Printf("\nYou used %s and did %d damage to the enemy!\n", usedItem.Name, enemyDamage)
			case "Give HP Shield":
				playerWizard.Health += usedItem.Price
				fmt.Printf("\nYou used %s that reduces incoming damage!\n", usedItem.Name)
			case "Give HP Helmet":
				playerWizard.Health += 20
				fmt.Printf("\nYou used %s and restored %d health!\n", usedItem.Name, 20)
			case "Give HP Torsal":
				playerWizard.Health += 30
				fmt.Printf("\nYou used %s and restored %d health!\n", usedItem.Name, 30)

			}

			inventory = append(inventory[:itemChoice-1], inventory[itemChoice:]...)
		default:
			fmt.Println("Incorrect choice.")
			continue
		}

		if enemyWizard.Health <= 0 {
			victoryWizard()
		}

		enemyAttack := WizardAttack{Name: "Enemy attacks", Damage: rand.Intn(10) + 10}
		playerWizard.Health -= enemyAttack.Damage
		fmt.Printf("Enemy did %d damage to you!\n", enemyAttack.Damage)

		if playerWizard.Health <= 0 {
			loseWizard()
		}
		Round++
	}
	fmt.Println("Game over.")
}

//relance une partie si on a gagné la précédente
func victoryWizard() {
	fmt.Println("You have won!")
	enemyWizard.Health = 75
	playerWizard.Health = 90

	//permet d'augmenter les HP de l'ennemi
	if ValueWizard == 0 {
		enemyWizard.Health = 100

	} else if ValueWizard == 1 {
		enemyWizard.Health = 125

	} else if ValueWizard == 2 {
		enemyWizard.Health = 150

		// 4 victoires = fin du jeu
	} else if ValueWizard == 3 {
		fmt.Println("\nYou've saved our world !")
		fmt.Println("Congratulations, you've just completed the game.")
		fmt.Println("\nr. Restart a new game")
		fmt.Println("q. quit game")
		fmt.Scanln(&WinWizard)
		switch WinWizard {
		case "r":
			Start()
		case "q":
			os.Exit(0)
		default:
			fmt.Println("Incorrect choice")
			victoryWizard()
		}
	}
	fmt.Println("Next game ?")
	fmt.Println("c. continue") //permet de recommencer une nouvelle partie
	fmt.Println("q. quit game")
	fmt.Scanln(&WinWizard)
	switch WinWizard {
	case "c":
		ValueWizard++
		inventory = []Object{}
		Health = 0
		Poison = 0
		Upgrade = 0
		Shieldd = 0
		Helmett = 0
		Torsal_armurr = 0
		stock = 0
		Coins = 100
		submenu_perso3()
	case "q":
		os.Exit(0)
	default:
		fmt.Println("Incorrect choice")
		victoryWizard()
	}
}

// Si l'ennemi gagne
func loseWizard() {
	fmt.Println("Enemy has won.")
	DefeatWizard++
	Start()
}
