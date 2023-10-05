package structure

import (
    "fmt"
    "math/rand"
    "time"
    "os"
)

//caractéristique du perso3
type Sorcier struct {
    Name   string //son nom
    Health int //ses points de vie de départ
}

//caractéristique des attaques du perso3
type SorcierAttack struct {
    Name   string //le nom
    Damage int //les dégats
}

//caractéristique de l'inventaire du perso3
type InventorySorcier struct {
    Name       string //nom de l'objet
    EffectType string //son type
    Value      int
}
var (enemySorcier = Sorcier{Name: "Enemy", Health: 75}//nom et points de vie du perso3
    playerSorcier = Sorcier{Name: "Reicros", Health: 90}//nom et points de vie de l'ennemi
    ValueDeSorcier int
    DefeatSorcier int
    WinDeSorcier string
)

//début du combat en tant que perso3
func GameSorcier(inventory []Object) {
    fmt.Print("\033[H\033[2J")
    rand.Seed(time.Now().UnixNano())

   
    if DefeatSorcier!=0 {
        playerSorcier.Health = 90
        enemySorcier.Health = 75
        DefeatSorcier = 0
        ValueDeSorcier = 0
    }
  
    attack1 := SorcierAttack{Name: "Fireball ", Damage: 20}
    attack2 := SorcierAttack{Name: "Lightning", Damage: 30}

    fmt.Println("Welcome to game!")
    
    TourDeCombat := 1

    for playerSorcier.Health > 0 && enemySorcier.Health > 0 {

        fmt.Printf("\nRound: %d\n", TourDeCombat)
        fmt.Printf("%s (HP: %d) vs %s (HP: %d)\n", playerSorcier.Name, playerSorcier.Health, enemySorcier.Name, enemySorcier.Health)
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

            var playerAttack SorcierAttack
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
        
            enemySorcier.Health -= enemyDamage
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
                playerSorcier.Health += 15
                fmt.Printf("\nYou used %s and restored %d health!\n", usedItem.Name, 15)
            case "Poison":
                enemySorcier.Health -= 15
                fmt.Printf("\nYou used %s and dealt %d damage to the enemy!\n", usedItem.Name, 15)
            case "Upgrade":
                playerAttack := attack1
                playerAttack.Damage = int(float64(playerAttack.Damage) * 1.2)
                enemyDamage := rand.Intn(playerAttack.Damage)
                enemySorcier.Health -= enemyDamage
                fmt.Printf("\nYou used %s and did %d damage to the enemy!\n", usedItem.Name, enemyDamage)
            case "Give HP Shield":
                playerSorcier.Health += usedItem.Price
                fmt.Printf("\nYou used %s that reduces incoming damage!\n", usedItem.Name)
            case "Give HP Helmet":
                playerSorcier.Health += 20
                fmt.Printf("\nYou used %s and restored %d health!\n", usedItem.Name, 20)
            case "Give HP Torsal":
                playerSorcier.Health += 30
                fmt.Printf("\nYou used %s and restored %d health!\n", usedItem.Name, 30)
            
            }
       
            
            inventory = append(inventory[:itemChoice-1], inventory[itemChoice:]...)
        default:
            fmt.Println("Incorrect choice.")
            continue
        }

        if enemySorcier.Health <= 0 {
            victorySorcier()
        }

        enemyAttack := SorcierAttack{Name: "Enemy attacks", Damage: rand.Intn(10) + 10}
        playerSorcier.Health -= enemyAttack.Damage
        fmt.Printf("Enemy did %d damage to you!\n", enemyAttack.Damage)
      
        if playerSorcier.Health <= 0 {
            loseSorcier()
        }
        TourDeCombat++
    }
    fmt.Println("Game over.")
}

//relance une partie si on a gagné la précédente
func victorySorcier() {
    fmt.Println("You have won!")
    enemySorcier.Health =75
    playerSorcier.Health=90
   
    if ValueDeSorcier == 0 {
        enemySorcier.Health =100
       
    } else if ValueDeSorcier == 1 {
        enemySorcier.Health =125
       
    }else if ValueDeSorcier == 2 {
        enemySorcier.Health = 150
    } else if ValueDeSorcier == 3{
        fmt.Println("\nYou've saved our world !")
        fmt.Println("Congratulations, you've just completed the game.")
    
    fmt.Println("\nr. Restart a new game")
    fmt.Println("q. quit game")
    fmt.Scanln(&WinDeSorcier)
    switch WinDeSorcier {
        case "r":
            
           Start()
        case "q":
            os.Exit(0)
        default:
            fmt.Println("Incorrect choice")
            victorySorcier()
    }
    }
    fmt.Println("Next game ?")
    fmt.Println("c. continue")
    fmt.Println("q. quit game")
    fmt.Scanln(&WinDeSorcier)
    switch WinDeSorcier {
        case "c":
            ValueDeSorcier++
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
            victorySorcier()
    }
}

func loseSorcier() {
    fmt.Println("Enemy has won.")
    DefeatSorcier++
    Start()
    //suite
}