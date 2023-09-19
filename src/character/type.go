package character

import "fmt"

type Perso struct {
	Name string
	Class string
	Level int
	HPmax int
	HPmin int
	Inventory int
	Money int
}

func Personne() {
    perso1 := []string{"nom1,", "classe1,", "niveau1,", "hpmax1,", "hpact1,", "inventaire1,", "argent1"}
    perso2 := []string{"nom2,", "classe2,", "niveau2,", "hpmax2,", "hpact2,", "inventaire2,", "argent2"}
    perso3 := []string{"nom3,", "classe3,", "niveau3,", "hpmax3,", "hpact3,", "inventaire3,", "argent3"}
    fmt.Println(perso1)
    fmt.Println(perso2)
    fmt.Println(perso3)
}