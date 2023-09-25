package structure

import (
	"fmt"
	"os"
)

type Object struct {
    Name string
    Effect string
    Price string
}


var choice string
var coins int = 100
var stock int = 0
var stockmax int = 5
var Health int = 0
var Poison int = 0
var Upgrade int = 0
var Shieldd int = 0

func Obj() {
    Health_potion := []string{"1. Health potion:", "Donne 15HP et coute 25 $"}
    Poison_potion := []string{"2. Poison potion:", "Enlève 15HP à l'adversaire et coute 25 $"}
    Upgrade_potion := []string{"3. Upgrade potion:", "Améliore tes attaques de 20% sur ton prochain tour et coute 50 $"}
    Shield := []string{"4. Shield:", "Réduit les dégats de 15% de la prochaine attaque subit et coute 40 $"}
    fmt.Println(Health_potion)
    fmt.Println(Poison_potion)
    fmt.Println(Upgrade_potion)
    fmt.Println(Shield)
    ChoixObj()
}
func ChoixObj() {
	//ChoixObj()
    fmt.Println("\ng. Go!")
    fmt.Println("i. Show inventaire")
	fmt.Println("r. Return to previous menu")
	fmt.Println("q. Quit game")
    fmt.Scanln(&choice)
	switch choice {
		case "1":
		Health_potion() 
		case "2":
		Poison_potion() 
		case "3":
		Upgrade_potion() 
		case "4":
		Shield()
        case "g":
       Game()
        
        case "i":
        Inventaire()
        case "r":
		ChoixPersonne()
		case "q":
		os.Exit(0)
		default:
		fmt.Println("Incorrect choice.")
		Obj()
	}
}
func Health_potion() {
    if stock == stockmax {
        fmt.Println("No more place in your inventory")
        Obj()
    } else if coins < 25 {
        fmt.Println("Not enough coins for this item")
        Obj()
    } else {
        coins = coins - 25
        stock = stock + 1
        Health = Health + 1
        fmt.Println("You have now",  coins, "gold coins and", stock, "item on", stockmax)
        Obj()
    }
}
func Poison_potion() {
    if stock == stockmax {
        fmt.Println("No more place in your inventory")
        Obj()
    } else if coins < 25 {
        fmt.Println("Not enough coins for this item")
        Obj()
    } else {
        coins = coins - 25
        stock = stock + 1
        Poison = Poison + 1
        fmt.Println("You have now",  coins, "gold coins and", stock, "item on", stockmax)
        Obj()
    }
}
func Upgrade_potion() {
    if stock == stockmax {
        fmt.Println("No more place in your inventory")
        Obj()
    } else if coins < 50 {
        fmt.Println("Not enough coins for this item")
        Obj()
    } else {
        coins = coins - 50
        stock = stock + 1
        Upgrade = Upgrade + 1
        fmt.Println("You have now",  coins, "gold coins and", stock, "item on", stockmax)
        Obj()
    }
}
func Shield() {
    if stock == stockmax {
        fmt.Println("No more place in your inventory")
        Obj()
    } else if coins < 40 {
        fmt.Println("Not enough coins for this item")
        Obj()
        } else {
        coins = coins - 40
        stock = stock + 1
        Shieldd = Shieldd + 1
        fmt.Println("You have now",  coins, "gold coins and", stock, "item on", stockmax)
        Obj()
    }
}
func Inventaire() {
    a := Health
    b := Poison
    c := Upgrade
    d := Shieldd
    fmt.Println("\nYou have:")
    for i := 0; i <= stock-1; i++ {
        if a != 0 {
            fmt.Println(a, "health potion")
            a = 0
        } else if b != 0 {
            fmt.Println(b, "poison potion")
            b = 0
        } else if c != 0 {
            fmt.Println(c, "upgrade potion")
            c = 0
        } else if d != 0 {
            fmt.Println(d, "shield")
            d = 0
        }
    }
    Obj()
}