package structure

import ("fmt"

)
var sommaire2 []string

type Perso struct {  //caractéristique d'un personnage
    Name string //son nom
    Class string //sa classe
    Level int //son niveau
    HPmax int //point de vie maximum
    HPact int //point de vie de départ
    Inventory int //inventaire
    Money int //pièces
}

//affiche les noms et classe des personnages
func ChoixPersonne() {
    sommaire1 := []string{"   Name:     ", "Class:"}
    sommaire2 = []string{"Level :  ", "HPmax :  ", "HPact :  "}
    Tank := []string{"1. Tenace   ", "   Tank"}
    Elfe := []string{"2. Chiro    ", "   Elfe"}
    Sorcier := []string{"3. Reicros  ", "Sorcier"}
    fmt.Println(sommaire1)
    fmt.Println(Tank)
    fmt.Println(Elfe)
    fmt.Println(Sorcier)
}

//affiche les infos complémentaires du perso1
func InfoTank() {
    TankInfo := []string{"lvl 1     ", "200      ", "100     "}
    fmt.Print("      ")
    fmt.Println(sommaire2)
    fmt.Print("      ")
    fmt.Println(TankInfo)
    fmt.Println("")
    AttackTank()
 
}

//affiche les infos complémentaires du perso2
func InfoElfe() {
    ElfeInfo := []string{"lvl 1     ", "140      ", "70       "}
    fmt.Println(sommaire2)
    fmt.Println(ElfeInfo)
    AttackElfe()
    
}

//affiche les infos complémentaires du perso3
func InfoSorcier() {
    SorcierInfo := []string{"lvl 1     ", "160      ", "80       "}
    fmt.Println(sommaire2)
    fmt.Println(SorcierInfo)
    AttackSorcier()
}


