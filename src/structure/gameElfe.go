package structure

import (
    "fmt"
    "math/rand"
    "time"
    "os"
)

//caractéristique du perso2
type Elfe struct {
    Name   string //son nom
    Health int //ses points de vie de départ
}

//caractéristique des attaques du perso2
type ElfeAttack struct {
    Name   string //le nom
    Damage int //les dégats
}

//caractéristique de l'inventaire du perso2
type InventoryElfe struct {
    Name       string //nom de l'objet
    EffectType string //son type
    Value      int 
}
var (enemyElfe = Tank{Name: "Enemy", Health: 75}  //nom et points de vie de l'ennemi
    playerElfe = Tank{Name: "Chiro", Health: 80} //nom et points de vie du perso2 
    ValueDeElfe int
    WinDeElfe string
)
//initialisation du combat en tant que perso2
func GameElfe(inventory []Object) {
    rand.Seed(time.Now().UnixNano())



   
    
    attack1 := ElfeAttack{Name: "Life thief ", Damage: 15}
    attack2 := ElfeAttack{Name: "Golden arrow ", Damage: 25}

    fmt.Println("Welcome to game!")

    TourDeCombat := 1

    for playerElfe.Health > 0 && enemyElfe.Health > 0 { //vérifie si les 2 ont de la vie
        fmt.Printf("\nRound: %d\n", TourDeCombat)
        fmt.Printf("%s (HP: %d) vs %s (HP: %d)\n", playerElfe.Name, playerElfe.Health, enemyElfe.Name, enemyElfe.Health)
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
    
                var playerAttack ElfeAttack
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

                enemyElfe.Health -= enemyDamage
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
            playerElfe.Health += usedItem.Price
            fmt.Printf("\nYou used %s and restored %d health!\n", usedItem.Name, 15)
        case "Poison":
            enemyElfe.Health -= usedItem.Price
            fmt.Printf("\nYou used %s and dealt %d damage to the enemy!\n", usedItem.Name, 15)
        case "Upgrade":
            playerAttack := attack1
            playerAttack.Damage = int(float64(playerAttack.Damage) * 1.2)
            enemyDamage := rand.Intn(playerAttack.Damage)
            enemyElfe.Health -= enemyDamage
            fmt.Printf("\nYou used %s and did %d damage to the enemy!\n", usedItem.Name, enemyDamage)
        case "Shield":
            playerElfe.Health += usedItem.Price
            fmt.Printf("\nYou used %s that reduces incoming damage!\n", usedItem.Name)
        }
        inventory = append(inventory[:itemChoice-1], inventory[itemChoice:]...)
        default:
            fmt.Println("Incorrect choice.")
            continue
        }

        if enemyElfe.Health <= 0 {
            victoryElfe()
        }

        EnemyAttack := ElfeAttack{Name: "Enemy attacks", Damage: rand.Intn(10) + 10}
        playerElfe.Health -= EnemyAttack.Damage
        fmt.Printf("Enemy did %d damage to you!\n", EnemyAttack.Damage)

        if playerElfe.Health <= 0 {
            loseElfe()
        }
        TourDeCombat++
    }
    fmt.Println("Game over.")
}

//relance une partie si on a gagné la précédente
func victoryElfe () {
    fmt.Println("You have won!")
    enemyElfe.Health =75
    playerTank.Health=80
    ValueDeElfe++
    if ValueDeElfe == 1 {
        enemyElfe.Health =100
       
    } else if ValueDeElfe == 2 {
        enemyElfe.Health =125
       
    }else if ValueDeElfe == 3 {
        enemyElfe.Health = 150
    } else if ValueDeElfe == 4{
        fmt.Println("\nYou've saved our world !")
        fmt.Println("Congratulations, you've just completed the game.")
    
    fmt.Println("\nr. Restart a new game")
    fmt.Println("q. quit game")
    fmt.Scanln(&WinDeElfe)
    switch WinDeElfe {
        case "r":
            
           Start()
        case "q":
            os.Exit(0)
        default:
            fmt.Println("Incorrect choice")
            victoryElfe()
    }
    }
    fmt.Println("Next game ?")
    fmt.Println("Yes. continue")
    fmt.Println("No. quit game")
    fmt.Scanln(&WinDeElfe)
    switch WinDeElfe {
        case "Yes":
            
           submenu_perso2()
        case "No":
            os.Exit(0)
        default:
            fmt.Println("Incorrect choice")
            victoryElfe()
    }
}

func loseElfe() {
    fmt.Println("Enemy has won.")
    Start()
    //suite
}