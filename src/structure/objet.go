package structure

import (
    "fmt"
    "os"
)

type Object struct { //caractéristique d'un objet
    Name   string //nom de l'objet
    Effect string //ce qu'il fait
    Price  int  //son prix
}

var choice string
var coins int = 100
var stock int = 0
var stockmax int = 5
var Health int = 0
var Poison int = 0
var Upgrade int = 0
var Shieldd int = 0

//création des objets
func Obj() {
    Health_potion := []string{"1. Health potion:", "Donne 15HP et coûte 25 $"}
    Poison_potion := []string{"2. Poison potion:", "Enlève 15HP à l'adversaire et coûte 25 $"}
    Upgrade_potion := []string{"3. Upgrade potion:", "Améliore tes attaques de 20%, sur ton prochain tour et coûte 50 $"}
    Shield := []string{"4. Shield:", "Réduit les dégâts de 15%, de la prochaine attaque subie et coûte 40 $"}
    fmt.Println(Health_potion)
    fmt.Println(Poison_potion)
    fmt.Println(Upgrade_potion)
    fmt.Println(Shield)
    ChoixObj()
}

//achats des objets et/ou début du jeu
func ChoixObj() {
    fmt.Println("\ng. Go!")
    fmt.Println("i. Show inventaire")
	fmt.Println("r. Return to previous menu")
	fmt.Println("q. Quit game")
    fmt.Scanln(&choice)
	switch choice {
		case "1":
		Health_potion()  //achat d'une potion de vie
		case "2":
		Poison_potion() //achat d'une potion de poison
		case "3":
		Upgrade_potion()  //achat d'une potion d'amélioration
		case "4":
		Shield() //achat d'un bouclier
        case "g":
      if ChoicePerso== 1{
        GameTank(inventory) //en tant que Tank (perso1)
      } else if ChoicePerso ==2 {
        GameElfe(inventory) //en tant qu'Elfe (perso2)
      } else if ChoicePerso == 3 {
        GameSorcier(inventory)  //en tant que Sorcier (perso3)
      }
        
        case "i":
        Inventaire() //affiche les objets achetés
        case "r":
		ChoixPersonne() //retour en arrière
		case "q":
		os.Exit(0) //Quitter
		default:
		fmt.Println("Incorrect choice.") //Erreur
		Obj()
	}
}


var inventory []Object


func GetInventory() []Object {
    return inventory
}

//permet d'ajouter à l'inventaire l'objet acheté
func AddToInventory(inventory []Object, item Object) []Object { 
    if len(inventory) < stockmax && coins >= item.Price { //vérifie que le stock ne soit pas plein et qu'il y ait assez de pièces pour acheter l'objet

        for i, invItem := range inventory { // vérifie si il y a quelque chose dans l'inventaire
            if invItem.Name == item.Name {  //si un objet se trouve dans l'inventaire
                inventory[i].Price += item.Price //augmenter l'inventaire
                coins -= item.Price //réduire les pièces
                fmt.Printf("You have now %d gold coins and %d %s in your inventory.\n", coins, inventory[i].Price, item.Name)
                return inventory
            }
        }

        inventory = append(inventory, item)
        coins -= item.Price
    
    } else if len(inventory) >= stockmax { //pour si l'inventaire est plein
        fmt.Println("No more place in your inventory")
    } else { //ou si il n'y a pas assez de pièces
        fmt.Println("Not enough coins for this item")
    }
    return inventory
}




//permet d'ajouter une potion de vie dans l'inventaire
func Health_potion() {
    if stock == stockmax {
        fmt.Println("No more place in your inventory")
        Obj()
    } else if coins < 25 {
        fmt.Println("Not enough coins for this item")
        Obj()
    } else {

    item := Object{Name: "Health potion", Effect: "Heal", Price: 25}
    inventory = AddToInventory(inventory, item)
   
        stock = stock + 1
        Health = Health + 1
        fmt.Println("You have now",  coins, "gold coins and", stock, "item on", stockmax)
        Obj()
    }
}

//permet d'ajouter une potion de poison dans l'inventaire
func Poison_potion() {
    if stock == stockmax {
        fmt.Println("No more place in your inventory")
        Obj()
    } else if coins < 25 {
        fmt.Println("Not enough coins for this item")
        Obj()
    } else {
        item := Object{Name: "Poison potion", Effect: "Poison", Price: 25}
        inventory = AddToInventory(inventory, item)
     
        stock = stock + 1
        Poison = Poison + 1
        fmt.Println("You have now",  coins, "gold coins and", stock, "item on", stockmax)
        Obj()
    }
}

//permet d'ajouter une potion d'amélioration dans l'inventaire
func Upgrade_potion() {
    if stock == stockmax {
        fmt.Println("No more place in your inventory")
        Obj()
    } else if coins < 50 {
        fmt.Println("Not enough coins for this item")
        Obj()
    } else {
        item := Object{Name: "Upgrade potion", Effect: "Upgrade", Price: 50}
        inventory = AddToInventory(inventory, item)
     
        stock = stock + 1
        Upgrade = Upgrade + 1
        fmt.Println("You have now",  coins, "gold coins and", stock, "item on", stockmax)
        Obj()
    }
}

//permet d'ajouter un bouclier dans l'inventaire
func Shield() {
    if stock == stockmax {
        fmt.Println("No more place in your inventory")
        Obj()
    } else if coins < 40 {
        fmt.Println("Not enough coins for this item")
        Obj()
        } else {

        item := Object{Name: "Shield", Effect: "Reduit Degat", Price: 40}
        inventory = AddToInventory(inventory, item)
     
        stock = stock + 1
        Shieldd = Shieldd + 1
        fmt.Println("You have now",  coins, "gold coins and", stock, "item on", stockmax)
        Obj()
    }
}
func Inventaire() {
    fmt.Println("\nYou have:")
    for i := 0; i <= stock-1; i++ {
        if Health != 0 {
            fmt.Println(Health, "health potion")
            Health = 0
        } else if Poison != 0 {
            fmt.Println(Poison, "poison potion")
            Poison = 0
        } else if Upgrade != 0 {
            fmt.Println(Upgrade, "upgrade potion")
            Upgrade = 0
        } else if Shieldd != 0 {
            fmt.Println(Shieldd, "shield")
            Shieldd = 0
        }
    }
    Obj()
}