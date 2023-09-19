package main

import (
	"fmt"
	"game/character"
)


func main() {
	conserve := character.Perso{
		"Pokemon",
		"Ranger",
		1,
		50,
		25,
		0,
		2,
	}
	
	for {
		fmt.Println("a. Info perso")
		fmt.Println("b. Inventory")
		fmt.Println("q. Quitter")

		var choice string
		fmt.Print("Enter letter of choice ")
		_, err := fmt.Scan(&choice)

		if err != nil {
		fmt.Println("Input Error. Please enter a letter (a, b or q).")
		continue
		}
		
		switch choice {
			case "a":
			fmt.Println("selected 'Character Information'.")
			fmt.Println("Name: ",getPerso(&conserve))
			case "b":
			fmt.Println("selected 'Inventory'.")
   
			case "q":
			fmt.Println("Goodbye!")
			return
			default:
			fmt.Println("Incorrect choice. Please select an option from the menu (a, b or q)")
		}
	}
}
func getPerso(p *character.Perso) string {
	return p.Name
}
func EditPerso(p character.Perso) string {
	return p.Name
}