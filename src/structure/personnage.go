package structure

import "fmt"

type Perso struct {
    Name string
    Class string
    Level int
    HPmax int
    HPact int
    Inventory int
    Money int
}

func Personne() {
    sommaire := []string{"Name :   ", "Class :  ", "Level :  ", "HPmax :  ", "HPact :  ", "Inventory :", "Money :"}
    perso1 := []string{"Tenace   ", "Tank     ", "Level 1  ", "200      ", "100      ", "0          ", "100 $  "}
    perso2 := []string{"Chiro    ", "Elfe     ", "Level 1  ", "140      ", "70       ", "0          ", "100 $  "}
    perso3 := []string{"Reicros  ", "Sorcier  ", "Level 1  ", "160      ", "80       ", "0          ", "100 $  "}
    fmt.Println(sommaire)
    fmt.Println(perso1)
    fmt.Println(perso2)
    fmt.Println(perso3)
}