package structure

import (
    "fmt"
    "os"
)

type Object struct { //caractéristique d'un objet
    Name   string //nom de l'objet
    Effect string //ce qu'il fait
    Price  int //son prix
    Sentence string //phrase en plus pour mettre tout ensemble
    Stock int //stock
}

var choice string
var Coins int = 100
var stock int = 0
var stockmax int = 5
var Health int = 0
var Poison int = 0
var Upgrade int = 0
var Shieldd int = 0
var Helmett int = 0
var Torsal_armurr int= 0

//création des objets
func Obj() {

    fmt.Println("\nI N V E N T O R Y")
    fmt.Printf("")
    fmt.Println("You can buy item only one time")
    fmt.Println("You have", Coins, "gold coins and", stock, "item on", stockmax)
	fmt.Print("\n")

    Health_potion := Object{Name: "1. Health potion :", Effect: "Give you 15HP and costs", Price: 25, Sentence: "$. Number of health potion :", Stock: 1}
    Poison_potion := Object{Name: "2. Poison potion :", Effect: "Make lose 15HP to the enemy and costs", Price: 25, Sentence: "$. Number of poison potion :", Stock: 1}
    Upgrade_potion := Object{Name: "3. Upgrade potion :", Effect: "Upgrade your next attack by 20%, one time and costs", Price: 40, Sentence: "$. Number of upgrade potion :", Stock: 1}
    Shield :=  Object{Name: "4. Shield :", Effect: "Reduces your damage by 15%, of the next attack suffered and costs", Price: 35, Sentence: "$. Number of shield :", Stock: 1}
    //Helmet := Object{Name: "5. Helmet :", Effect: "Give you 20HP and costs", Price: 40, Sentence: "$. Number of helmet :", Stock: 1}
    //Torsal_armur := Object{Name: "6. Torsal armur :", Effect: "Give you 30HP and costs", Price: 60, Sentence: "$. Number of torsal armur :", Stock: 1}
    
    
    if Health == 0 {
        fmt.Println(Health_potion)
    } 
    if Poison == 0 {
        fmt.Println(Poison_potion)
    } 
    if Upgrade == 0 {
        fmt.Println(Upgrade_potion)
    }
    if Shieldd == 0 {
        fmt.Println(Shield)
    }
    //if Helmett == 0 {
        //fmt.Println(Helmet)
    //}
    //if Torsal_armurr == 0 {
        //fmt.Println(Torsal_armur)
    //}
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
            if Health == 0 {
                Health_potion() //achat d'une potion de vie
            } else {
                fmt.Println("You've already bought this item")
                Obj()
            }
		case "2":
            if Poison == 0 {
                Poison_potion() //achat d'une potion de poison
            } else {
                fmt.Println("You've already bought this item")
                Obj()
            }
		case "3":
            if Upgrade == 0 {
                Upgrade_potion() //achat d'une potion d'amélioration
            } else {
                fmt.Println("You've already bought this item")
                Obj()
            }
		case "4":
            if Shieldd == 0 {
                Shield() //achat d'un bouclier
            } else {
                fmt.Println("You've already bought this item")
                Obj()
            }
        case "g": //lancement du jeu
            if ChoicePerso == 1{
                GameTank(inventory) //en tant que Tank (perso1)
      }     else if ChoicePerso == 2 {
                GameElfe(inventory) //en tant qu'Elfe (perso2)
      }     else if ChoicePerso == 3 {
               GameSorcier(inventory) //en tant que Sorcier (perso3)
      }

        case "i":
            Inventaire() //affiche les objets achetés
        case "r":
		    ChoixPersonne() //retour en arrière
		case "q":
		    os.Exit(0) //Quitter
		default:
		    fmt.Println("Incorrect choice.")
		    Obj() //si le choix fait est faux
	}
}

var inventory []Object 

func GetInventory() []Object {
    return inventory
}

//permet d'ajouter à l'inventaire l'objet acheté
func AddToInventory(inventory []Object, item Object) []Object{
    if len(inventory) < stockmax && Coins >= item.Price { //vérifie que le stock ne soit pas plein et qu'il y ait assez de pièces pour acheter l'objet
        for i, invItem := range inventory { // vérifie si il y a quelque chose dans l'inventaire
            if invItem.Name == item.Name { //si un objet se trouve dans l'inventaire
                inventory[i].Price += item.Price //augmenter l'inventaire
                Coins -= item.Price //réduire les pièces
                return inventory
            }
        }
        inventory = append(inventory, item)
        Coins -= item.Price
        fmt.Printf("You have now %d gold coins and %s in your inventory.\n", Coins, item.Name)
    
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
    } else if Coins < 25 {
        fmt.Println("Not enough coins for this item")
    Obj() 
    } else {

    item := Object{Name: "Health potion", Effect: "Heal", Price: 25, Stock: 1}
    inventory = AddToInventory(inventory, item)

        stock = stock + 1
        Health = Health + 1

    Obj() 
    }
}

//permet d'ajouter une potion de poison dans l'inventaire
func Poison_potion() {
    if stock == stockmax {
        fmt.Println("No more place in your inventory")
    Obj() 
    } else if Coins < 25 {
        fmt.Println("Not enough coins for this item")
    Obj() 
    } else {
        item := Object{Name: "Poison potion", Effect: "Poison", Price: 25, Stock: 1}
        inventory = AddToInventory(inventory, item)
     
        stock = stock + 1
        Poison = Poison + 1

    Obj() 
    }
}

//permet d'ajouter une potion d'amélioration dans l'inventaire
func Upgrade_potion() {
    if stock == stockmax {
        fmt.Println("No more place in your inventory")
    Obj() 
    } else if Coins < 50 {
        fmt.Println("Not enough coins for this item")
    Obj() 
    } else {
        item := Object{Name: "Upgrade potion", Effect: "Upgrade", Price: 50, Stock: 1}
        inventory = AddToInventory(inventory, item)
     
        stock = stock + 1
        Upgrade = Upgrade + 1
        
    Obj() 
    }
}

//permet d'ajouter un bouclier dans l'inventaire
func Shield() {
    if stock == stockmax {
        fmt.Println("No more place in your inventory")
    Obj() 
    } else if Coins < 40 {
        fmt.Println("Not enough coins for this item")
    Obj() 
        } else {

        item := Object{Name: "Helmet", Effect: "Give HP", Price: 40, Stock: 1}
        inventory = AddToInventory(inventory, item)
     
        stock = stock + 1
        Shieldd = Shieldd + 1
        
    Obj() 
    }
}

//func Helmet() {
    //if stock == stockmax {
        //fmt.Println("No more place in your inventory")
    //Obj() 
    //} else if Coins < 40 {
        //fmt.Println("Not enough coins for this item")
    //Obj() 
        //} else {
        //item := Object{Name: "Shield", Effect: "Reduit Degat", Price: 40, Stock: 1}
        //inventory = AddToInventory(inventory, item)
     
        //stock = stock + 1
        //Helmett = Helmett + 1  
    //Obj() 
    //}
//}

//func Torsal_armur() {
    //if stock == stockmax {
        //fmt.Println("No more place in your inventory")
    //Obj() 
    //} else if Coins < 60 {
        //fmt.Println("Not enough coins for this item")
    //Obj() 
        //} else {
        //item := Object{Name: "Torsal armur", Effect: "Give HP", Price: 60, Stock: 1}
        //inventory = AddToInventory(inventory, item)
     
        //stock = stock + 1
        //Torsal_armurr = Torsal_armurr + 1
        
    //Obj() 
    //}
//}

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