package structure

import "fmt"

type Attaque struct {
    Name string
    Damage int
    Précision int
}

func Attack() {
    Tenace1 := []string{"Tremblement de terre", "25", "100"}
    Tenace2 := []string{"Coup de poing", "40", "55"}
    Chiro1 := []string{"Volvi", "15", "85"} //l'adversaire perd 15HP et Chiro gagne 15HP
    Chiro2 := []string{"Flèche d'or", "25", "80"}
    Reicros1 := []string{"Boule de feu", "20", "75"}
    Reicros2 := []string{"Jet de foudre", "30", "90"}
    fmt.Println(Tenace1)
    fmt.Println(Tenace2)
    fmt.Println(Chiro1)
    fmt.Println(Chiro2)
    fmt.Println(Reicros1)
    fmt.Println(Reicros2)
}