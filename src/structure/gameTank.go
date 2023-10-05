package structure

import (
    "fmt"
    "math/rand"
    "time"
    "os"
)

//caractéristique du perso1
type Tank struct {
    Name   string //son nom
    Health int //ses points de vie de départ
}

//caractéristique des attaques du perso1
type TankAttack struct {
    Name   string //le nom
    Damage int //les dégats
}

//caractéristique de l'inventaire du perso1
type InventoryTank struct {
    Name       string //nom de l'objet
    EffectType string //son type
    Value      int
}
var (enemyTank = Tank{Name: "Enemy", Health: 75}//nom et points de vie de l'ennemi
    playerTank = Tank{Name: "Tenace", Health: 100} //nom et points de vie du perso1
    ValueDeTank int
    DefeatTank int
    WinDeTank string
)


//début du combat en tant que perso1
func GameTank(inventory []Object) {
    fmt.Print("\033[H\033[2J")
    rand.Seed(time.Now().UnixNano())

  
    if DefeatTank!=0 {
        playerTank.Health = 80
        enemyTank.Health = 75
        DefeatTank = 0
        ValueDeTank = 0
    }
   

 

    attack1 := TankAttack{Name: "Earthquake ", Damage: 25}
    attack2 := TankAttack{Name: "Rock punch", Damage: 40}

    fmt.Println("Welcome to game!")

    TourDeCombat := 1

    for playerTank.Health > 0 && enemyTank.Health > 0 {

        fmt.Printf("\nRound: %d\n", TourDeCombat)
        fmt.Printf("%s (HP: %d) vs %s (HP: %d)\n", playerTank.Name, playerTank.Health, enemyTank.Name, enemyTank.Health)
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

            var playerAttack TankAttack
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

            enemyDamage = playerAttack.Damage
            enemyTank.Health -= enemyDamage
            fmt.Printf("\nYou did %d damage to enemy!\n", enemyDamage)

        case 2: //2 = utiliser un objet
            //a revoir
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
                playerTank.Health += 15
                fmt.Printf("\nYou used %s and restored %d health!\n", usedItem.Name, 15)
            case "Poison":
                enemyTank.Health -= 15
                fmt.Printf("\nYou used %s and dealt %d damage to the enemy!\n", usedItem.Name, 15)
            case "Upgrade":
                playerAttack := attack1
                playerAttack.Damage = int(float64(playerAttack.Damage) * 1.2)
                enemyDamage := playerAttack.Damage
                enemyTank.Health -= enemyDamage
                fmt.Printf("\nYou used %s and did %d damage to the enemy!\n", usedItem.Name, enemyDamage)
            case "Give HP Shield":
                playerTank.Health += usedItem.Price
                fmt.Printf("\nYou used %s that reduces incoming damage!\n", usedItem.Name)
                  case "Give HP Helmet":
                playerTank.Health += 20
                fmt.Printf("\nYou used %s and restored %d health!\n", usedItem.Name, 20)
            case "Give HP Torsal":
                playerTank.Health += 30
                fmt.Printf("\nYou used %s and restored %d health!\n", usedItem.Name, 30)
            }

            inventory = append(inventory[:itemChoice-1], inventory[itemChoice:]...)
        
        default:
            fmt.Println("Incorrect choice.")
            continue
        }
        if enemyTank.Health <= 0 {
            victoryTank()
        }

        enemyAttack := TankAttack{Name: "Enemy attacks", Damage: rand.Intn(10) + 10}
        playerTank.Health -= enemyAttack.Damage
        fmt.Printf("Enemy did %d damage to you!\n", enemyAttack.Damage)

        if playerTank.Health <= 0 {
            loseTank()
        }
        TourDeCombat++
    }
    fmt.Println("Game over.")
}
//relance une partie si on a gagné la précédente
func victoryTank() {
    fmt.Println("You have won!")
    enemyTank.Health =75
    playerTank.Health=100
   
    if ValueDeTank == 0 {
        enemyTank.Health =100
       
    } else if ValueDeTank == 1 {
        enemyTank.Health =125
       
    }else if ValueDeTank == 2 {
        enemyTank.Health = 150
    } else if ValueDeTank == 3{
        fmt.Println("\nYou've saved our world !")
        fmt.Println("Congratulations, you've just completed the game.")
    
    fmt.Println("\nr. Restart a new game")
    fmt.Println("q. quit game")
    fmt.Scanln(&WinDeTank)
    switch WinDeTank {
        case "r":
            
           Start()
        case "q":
            os.Exit(0)
        default:
            fmt.Println("Incorrect choice")
            victoryTank()
    }
    }
    fmt.Println("Next game ?")
    fmt.Println("c. continue")
    fmt.Println("q. quit game")
    fmt.Scanln(&WinDeTank)
    switch WinDeTank {
        case "c":
            ValueDeTank++
            inventory = []Object{}
            Health = 0
            Poison = 0
            Upgrade = 0
            Shieldd = 0
            Helmett = 0
            Torsal_armurr = 0
            stock = 0
            Coins = 100
           submenu_perso1()
        case "q":
            os.Exit(0)
        default:
            fmt.Println("Incorrect choice")
            victoryTank()
    }
}

func loseTank() {
    fmt.Println("Enemy has won.")
    DefeatTank++
    Start()
    //suite
}