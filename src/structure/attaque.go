package structure

import "fmt"

//caractéristique d'une attaque
type Attaque struct {
	Name      string //son nom
	Damage    int    //les dégats qu'elle inflige
	Précision int    //sa prédision (not used)
}

//affiche les caractéristique des attaques du perso1
func AttackTank() {
	sommaireattack := []string{"Attack's name:        ", "Damage: ", "Precision: "}
	Tenace1 := []string{"Earthquake            ", "25      ", "100        "}
	Tenace2 := []string{"Rock punch            ", "40      ", "55         "}
	fmt.Println(sommaireattack)
	fmt.Println(Tenace1)
	fmt.Println(Tenace2)
}

//affiche les caractéristique des attaques du perso2
func AttackElfe() {
	sommaireattack := []string{"Attack's name:        ", "Damage: ", "Precision: "}
	Chiro1 := []string{"Life thief            ", "15      ", "85         "}
	Chiro2 := []string{"Golden arrow          ", "25      ", "80         "}
	fmt.Println(sommaireattack)
	fmt.Println(Chiro1)
	fmt.Println(Chiro2)
}

//affiche les caractéristique des attaques du perso3
func AttackWizard() {
	sommaireattack := []string{"Attack's name:        ", "Damage: ", "Precision: "}
	Reicros1 := []string{"Fireball              ", "20      ", "75         "}
	Reicros2 := []string{"Lightning             ", "30      ", "90         "}
	fmt.Println(sommaireattack)
	fmt.Println(Reicros1)
	fmt.Println(Reicros2)
}
