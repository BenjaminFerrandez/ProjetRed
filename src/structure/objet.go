package structure

import (
    "fmt"
    "os"
)

type Object struct {
    Name   string
    Effect string
    Price  int 
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
    Health_potion := []string{"1. Health potion:", "Donne 15HP et coûte 25 $"}
    Poison_potion := []string{"2. Poison potion:", "Enlève 15HP à l'adversaire et coûte 25 $"}
    Upgrade_potion := []string{"3. Upgrade potion:", "Améliore tes attaques de 20% sur ton prochain tour et coûte 50 $"}
    Shield := []string{"4. Shield:", "Réduit les dégâts de 15% de la prochaine attaque subie et coûte 40 $"}
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
      if ChoicePerso== 1{
        GameTank()
      } else if ChoicePerso ==2 {
        GameElfe()
      } else if ChoicePerso == 3 {
        GameSorcier()
      }
        
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
var inventory []Object 


func AddToInventory(item Object) {
    
    if len(inventory) < stockmax && coins >= item.Price {
  
        for i, invItem := range inventory {
            if invItem.Name == item.Name {
                inventory[i].Price += item.Price
                coins -= item.Price
                fmt.Printf("You have now %d gold coins and %d %s(s) in your inventory.\n", coins, inventory[i].Price, item.Name)
                return
            }
        }
 
        inventory = append(inventory, item)
        coins -= item.Price
        fmt.Printf("You have now %d gold coins and %s in your inventory.\n", coins, item.Name)
    } else if len(inventory) >= stockmax {
        fmt.Println("No more place in your inventory")
    } else {
        fmt.Println("Not enough coins for this item")
    }
}


func ShowInventory() {
    fmt.Println("\nYour inventory:")
    for _, invItem := range inventory {
        fmt.Printf("%s (%d $)\n", invItem.Name, invItem.Price)
    }
}


func Health_potion() {
    item := Object{Name: "Health potion", Effect: "Heal", Price: 25}
   
    AddToInventory(item)
    Obj()
}
func Poison_potion() {
    item := Object{Name: "Poison potion", Effect: "Poison", Price: 25}
  
    AddToInventory(item)
    Obj()
}
func Upgrade_potion() {
    item := Object{Name: "Upgrade potion", Effect: "Upgrade", Price: 50}
 
    AddToInventory(item)
    Obj()
}
func Shield() {
    item := Object{Name: "Shield", Effect: "Shield", Price: 40}
  
    AddToInventory(item)
    Obj()
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