package structure

import (
    "fmt"
    "math/rand"
    "time"
)

type Elfe struct {
    Name   string
    Health int
}

type ElfeAttack struct {
    Name   string
    Damage int
}

type InventoryElfe struct {
    Name       string
    EffectType string
    Value      int
}

func GameElfe() {
    rand.Seed(time.Now().UnixNano())

    player := Elfe{Name: "Chiro", Health: 100}
    enemy := Elfe{Name: "Enemy", Health: 100}

    attack1 := ElfeAttack{Name: "Volvi", Damage: 15}
    attack2 := ElfeAttack{Name: "Flèche d'or ", Damage: 25}


    inventory := []InventoryElfe{
        {Name: "Healph potion", EffectType: "Heal", Value: 15},
        {Name: "Poison potion", EffectType: "Poison", Value: 15},
        {Name: "Upgrade potion", EffectType: "DamageBoost", Value: 20},
        {Name: "Shield", EffectType: "Shield", Value: 15},
    }

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
        case 1:
            fmt.Println("Choose attack:")
            fmt.Printf("1. %s (Damage: %d)\n", attack1.Name, attack1.Damage)
            fmt.Printf("2. %s (Damage: %d)\n", attack2.Name, attack2.Damage)

            var attackChoice int
            fmt.Print("Enter number of attack: ")
            fmt.Scanln(&attackChoice)

            var playerAttack ElfeAttack
            switch attackChoice {
            case 1:
                playerAttack = attack1
            case 2:
                playerAttack = attack2
            default:
                fmt.Println("Incorrect choice of attack.")
                continue
            }

            var enemyDamage int
            if attackChoice == 2 {
                if rand.Float64() < 0.6 {
                    enemyDamage = 40
                } else {
                    enemyDamage = 0
                }
            } else {
                enemyDamage = rand.Intn(playerAttack.Damage)
            }
            enemy.Health -= enemyDamage
            fmt.Printf("\nYou did %d damage to enemy!\n", enemyDamage)

        case 2:
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
            switch usedItem.EffectType {
            case "Heal":
                player.Health += usedItem.Value
                fmt.Printf("\nYou used %s and restored %d health!\n", usedItem.Name, usedItem.Value)
            case "Poison":
                enemy.Health -= usedItem.Value
                fmt.Printf("\nYou used %s and dealt %d damage to the enemy!\n", usedItem.Name, usedItem.Value)
            case "DamageBoost":
              
                playerAttack := attack1
                playerAttack.Damage = int(float64(playerAttack.Damage) * 1.2)
                enemyDamage := rand.Intn(playerAttack.Damage)
                enemy.Health -= enemyDamage
                fmt.Printf("\nYou used %s and did %d damage to the enemy!\n", usedItem.Name, enemyDamage)
            case "Shield":

                player.Health += usedItem.Value
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

        enemyAttack := ElfeAttack{Name: "Enemy attacks", Damage: rand.Intn(10) + 10}
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