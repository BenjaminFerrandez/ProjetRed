package structure

import "fmt"

type Attaque struct {
    Name string
    Damage int
    Précision int
}

func AttackTank() {
    sommaireattack := []string{"Attack's name:        ", "Damage: ", "Precision: "}
    Tenace1 :=        []string{"Tremblement de terre  ", "25      ", "100        "}
    Tenace2 :=        []string{"Coup de poing         ", "40      ", "55         "}
    fmt.Println(sommaireattack)
    fmt.Println(Tenace1)
    fmt.Println(Tenace2)
}
func AttackElfe() {
    sommaireattack := []string{"Attack's name:        ", "Damage: ", "Precision: "}
    Chiro1 :=         []string{"Volvi                 ", "15      ", "85         "} 
    Chiro2 :=         []string{"Flèche d'or           ", "25      ", "80         "}
    fmt.Println(sommaireattack)
    fmt.Println(Chiro1)
    fmt.Println(Chiro2)
}
func AttackSorcier() {
    sommaireattack := []string{"Attack's name:        ", "Damage: ", "Precision: "}
    Reicros1 :=       []string{"Boule de feu          ", "20      ", "75         "}
    Reicros2 :=       []string{"Jet de foudre         ", "30      ", "90         "}
    fmt.Println(sommaireattack)
    fmt.Println(Reicros1)
    fmt.Println(Reicros2)
}
