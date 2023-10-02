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

//initialisation du combat en tant que perso2
func GameElfe(inventory []Object) {
    rand.Seed(time.Now().UnixNano())

    player := Elfe{Name: "Chiro", Health: 80} //nom et points de vie du perso2
    Enemy := Enemy{Name: "Enemy", Health: 75} //nom et points de vie de l'ennemi

    for ValueDeJeu > 1 { //ajoute 25HP si l'ennemi a déja été battu
        Enemy.Health += 25
        
    }
    
    attack1 := ElfeAttack{Name: "Life thief ", Damage: 15}
    attack2 := ElfeAttack{Name: "Golden arrow ", Damage: 25}

    fmt.Println("Welcome to game!")

    TourDeCombat := 1

    for player.Health > 0 && Enemy.Health > 0 { //vérifie si les 2 ont de la vie
        fmt.Printf("\nRound: %d\n", TourDeCombat)
        fmt.Printf("%s (HP: %d) vs %s (HP: %d)\n", player.Name, player.Health, Enemy.Name, Enemy.Health)
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

                Enemy.Health -= enemyDamage
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
            player.Health += usedItem.Price
            fmt.Printf("\nYou used %s and restored %d health!\n", usedItem.Name, 15)
        case "Poison":
            Enemy.Health -= usedItem.Price
            fmt.Printf("\nYou used %s and dealt %d damage to the enemy!\n", usedItem.Name, 15)
        case "Upgrade":
            playerAttack := attack1
            playerAttack.Damage = int(float64(playerAttack.Damage) * 1.2)
            enemyDamage := rand.Intn(playerAttack.Damage)
            Enemy.Health -= enemyDamage
            fmt.Printf("\nYou used %s and did %d damage to the enemy!\n", usedItem.Name, enemyDamage)
        case "Shield":
            player.Health += usedItem.Price
            fmt.Printf("\nYou used %s that reduces incoming damage!\n", usedItem.Name)
        }
        inventory = append(inventory[:itemChoice-1], inventory[itemChoice:]...)
        default:
            fmt.Println("Incorrect choice.")
            continue
        }

        if Enemy.Health <= 0 {
            victoryElfe()
        }

        EnemyAttack := ElfeAttack{Name: "Enemy attacks", Damage: rand.Intn(10) + 10}
        player.Health -= EnemyAttack.Damage
        fmt.Printf("Enemy did %d damage to you!\n", EnemyAttack.Damage)

        if player.Health <= 0 {
            loseElfe()
        }
        TourDeCombat++
    }
    fmt.Println("Game over.")
}

//relance une partie si on a gagné la précédente
func victoryElfe () {
    fmt.Println("You have won!")
    fmt.Println("Next game ?")
    fmt.Println("Yes. continue")
    fmt.Println("No. quit game")
    fmt.Scanln(&win)
    switch win {
        case "Yes":
            ValueDeJeu += 1 
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
    //suite
}