package structure

import (
    "fmt"
    "os"
)

type Object struct { //caractéristique d'un objet
    Name   string //nom de l'objet
    Effect string //ce qu'il fait
    Price  int  //son prix
    Sentence string
    Stock int
}

var choice string
var coins int = 100 //on démarre avec 100 pièces,
    var stock int = 0 // 0 objet
    var stockmax int = 5 // et on peut avoir 5 objet au maximum
var Health int = 0
var Poison int = 0
var Upgrade int = 0
var Shieldd int = 0

//création des objets
func Obj() {
	
    fmt.Println("\nI N V E N T O R Y")
    fmt.Printf("")
    fmt.Println("You can buy item only one time")
    fmt.Println("You have", coins, "gold coins and", stock, "item on", stockmax)
	fmt.Print("\n")

    Health_potion := Object{Name: "1. Health potion :", Effect: "Give you 15HP and costs", Price: 25, Sentence: "$. Number of health potion :", Stock: 1}
    Poison_potion := Object{Name: "2. Poison potion :", Effect: "Enlève 15HP à l'adversaire et coûte", Price: 25, Sentence: "$. Number of poison potion :", Stock: 1}
    Upgrade_potion := Object{Name: "3. Upgrade potion :", Effect: "Améliore tes attaques de 20%, sur ton prochain tour et coûte", Price: 50, Sentence: "$. Number of upgrade potion :", Stock: 1}
    Shield :=  Object{Name: "4. Shield :", Effect: "Réduit tes dégats de 15%, de la prochaine attaque subie et coûte", Price: 40, Sentence: "$. Number of shield :", Stock: 1}
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
                Health_potion()  //achat d'une potion de vie
            } else   {
                fmt.Println("\nVous avez deja acheter potion de vie")
            Obj() 
            }
		case "2":
            if Poison == 0 {
                Poison_potion() //achat d'une potion de poison
            } else {
                fmt.Println("\nVous avez deja acheter poison potion")
            Obj() 
            }
		
		case "3":
            if Upgrade == 0 {
                Upgrade_potion()  //achat d'une potion d'amélioration
              
            } else {
                fmt.Println("\nVous avez deja acheter Upgrade potion")
            Obj() 
            }
	
		case "4":
            if Shieldd == 0 {
                Shield() //achat d'un bouclier
            } else {
                  fmt.Println("\nVous avez deja acheter shield")
            Obj() 
            }
	
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
    } else if coins < 25 {
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
    } else if coins < 50 {
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
    } else if coins < 40 {
        fmt.Println("Not enough coins for this item")
    Obj() 
        } else {

        item := Object{Name: "Shield", Effect: "Reduit Degat", Price: 40, Stock: 1}
        inventory = AddToInventory(inventory, item)
     
        stock = stock + 1
        Shieldd = Shieldd + 1
        
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