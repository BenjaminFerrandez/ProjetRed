package structure

import "fmt"

var sommaire2 []string

type Perso struct {
    Name string
    Class string
    Level int
    HPmax int
    HPact int
    Inventory int
    Money int
}

func ChoixPersonne() {
    sommaire1 := []string{"   Name:     ", "Class:"}
    sommaire2 = []string{"Level :  ", "HPmax :  ", "HPact :  ", "Inventory :", "Money :"}
    Tank := []string{"1. Tenace   ", "   Tank"}
    Elfe := []string{"2. Chiro    ", "   Elfe"}
    Sorcier := []string{"3. Reicros  ", "Sorcier"}
    fmt.Println(sommaire1)
    fmt.Println(Tank)
    fmt.Println(Elfe)
    fmt.Println(Sorcier)
}
func InitTank() {
    TankInfo := []string{"lvl 1     ", "200      ", "100      ", "0          ", "100 $ "}
    fmt.Println(sommaire2)
    fmt.Println(TankInfo)
 
}
func InitElfe() {
    ElfeInfo := []string{"lvl 1     ", "140      ", "70       ", "0          ", "100 $ "}
    fmt.Println(sommaire2)
    fmt.Println(ElfeInfo)
    
}
func InitSorcier() {
    SorcierInfo := []string{"lvl 1     ", "160      ", "80       ", "0          ", "100 $ "}
    fmt.Println(sommaire2)
    fmt.Println(SorcierInfo)
}