package structure

import (
    "fmt"
    "math/rand"
    "time"
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

//début du combat en tant que perso3
func GameSorcier(inventory []Object) {
    rand.Seed(time.Now().UnixNano())

    player := Sorcier{Name: "Reicros", Health: 100}
    enemy := Sorcier{Name: "Enemy", Health: 100}

    attack1 := SorcierAttack{Name: "Boule de feu", Damage: 25}
    attack2 := SorcierAttack{Name: "Jet de foudre", Damage: 30}

    fmt.Println("Welcome to game!")

    TourDeCombat := 1

    for player.Health > 0 && enemy.Health > 0 {

        fmt.Printf("\nTour de combat: %d\n", TourDeCombat)
        fmt.Printf("%s (HP: %d) vs %s (HP: %d)\n", player.Name, player.Health, enemy.Name, enemy.Health)
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
            switch attackChoice {
            case 1: //1.1 = attaque 1
                playerAttack = attack1
            case 2: //1.2 = attaque 2
                playerAttack = attack2
            default:
                fmt.Println("Incorrect choice of attack.")
                continue
            }

            enemyDamage := rand.Intn(playerAttack.Damage)
            enemy.Health -= enemyDamage
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
                player.Health += usedItem.Price
                fmt.Printf("\nYou used %s and restored %d health!\n", usedItem.Name, usedItem.Price)
            case "Poison":
                enemy.Health -= usedItem.Price
                fmt.Printf("\nYou used %s and dealt %d damage to the enemy!\n", usedItem.Name, usedItem.Price)
            case "Upgrade":
              
            case "Shield":
                player.Health += usedItem.Price
                fmt.Printf("\nYou used %s that reduces incoming damage!\n", usedItem.Name)
            }

            inventory = append(inventory[:itemChoice-1], inventory[itemChoice:]...)

        default:
            fmt.Println("Incorrect choice.")
            continue
        }

        if enemy.Health <= 0 {
            fmt.Println("You have won!")
            break
        }
enemyAttack := SorcierAttack{Name: "Enemy attacks", Damage: rand.Intn(10) + 10}
        player.Health -= enemyAttack.Damage
        fmt.Printf("Enemy did %d damage to you!\n", enemyAttack.Damage)

        if player.Health <= 0 {
            fmt.Println("Enemy has won.")
            break
        }

        TourDeCombat++
    }

    fmt.Println("Game over.")
}