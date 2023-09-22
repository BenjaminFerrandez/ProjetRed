package structure

import "fmt"

type Object struct {
    Name string
    Effect string
    Price string
}

func Obj() {
    Health_potion := []string{"Health potion:", "Donne 15HP", "Coute 25 $"}
    Poison_potion := []string{"Poison potion:", "Enlève 15HP à l'adversaire", "Coute 25 $"}
    Upgrade_potion := []string{"Upgrade potion:", "Améliore tes attaques de 20% sur ton prochain tour", "Coute 50 $"}
    Shield := []string{"Shield:", "Réduit les dégats de 15% de la prochaine attaque subit", "Coute 40 $"}
    fmt.Println(Health_potion)
    fmt.Println(Poison_potion)
    fmt.Println(Upgrade_potion)
    fmt.Println(Shield)
}