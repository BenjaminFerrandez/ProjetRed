package structure

import "fmt"

type Object struct {
    Name string
    Effect string
    Price string
}

func Obj() {
    Health_potion := []string{"Health potion", "Donne XXHP", "Coute XX $"}
    Poison_potion := []string{"Poison potion", "Enlève XXHP à l'adversaire", "Coute XX $"}
    Upgrade_potion := []string{"Upgrade potion", "améliore tes attaques de XX% sur ton prochain tour", "Coute XX $"}
    Shield := []string{"Shield", "Réduit les dégats de XX% de la prochaine attaque subit", "Coute XX $"}
    fmt.Println(Health_potion)
    fmt.Println(Poison_potion)
    fmt.Println(Upgrade_potion)
    fmt.Println(Shield)
}