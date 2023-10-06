package structure

import "fmt"

var sommaire2 []string

type Perso struct { //caractéristique d'un personnage
	Name      string //son nom
	Class     string //sa classe
	HPmax     int    //point de vie maximum
	HPact     int    //point de vie de départ
	Inventory int    //inventaire
	Money     int    //pièces
}

//affiche les noms et classe des personnages
func ChoiceChar() {
	sommaire1 := []string{"   Name:     ", "Class:"}
	sommaire2 = []string{"HPmax :  ", "HPact :  "}
	Tank := []string{"1. Tenace   ", "   Tank"}
	Elfe := []string{"2. Chiro    ", "   Elfe"}
	Wizard := []string{"3. Reicros  ", " Wizard"}
	fmt.Println(sommaire1)
	fmt.Println(Tank)
	fmt.Println(Elfe)
	fmt.Println(Wizard)
}

//affiche les infos complémentaires du perso1
func InfoTank() {
	TankInfo := []string{"200      ", "100     "}
	fmt.Print("      ")
	fmt.Println(sommaire2)
	fmt.Print("      ")
	fmt.Println(TankInfo)
	fmt.Println("")
	AttackTank()

}

//affiche les infos complémentaires du perso2
func InfoElfe() {
	ElfeInfo := []string{"160      ", "80       "}
	fmt.Print("      ")
	fmt.Println(sommaire2)
	fmt.Print("      ")
	fmt.Println(ElfeInfo)
	fmt.Print("")
	AttackElfe()

}

//affiche les infos complémentaires du perso3
func InfoWizard() {
	WizardInfo := []string{"180      ", "90       "}
	fmt.Print("      ")
	fmt.Println(sommaire2)
	fmt.Print("      ")
	fmt.Println(WizardInfo)
	fmt.Print("")
	AttackWizard()
}
