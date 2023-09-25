package main

import ("fmt"
    "game/structure"
)

func main() {
    for {
        fmt.Println("\nBienvenue")
        fmt.Println("Click 's' to start or 'q' to quit")
        var choice string
        fmt.Scanln(&choice)
        switch choice {
        	case "s":
            fmt.Println("Goodluck and enjoy !")
            structure.Start()
       		case "q":
            fmt.Println("Goodbye!")
            return
       		default:
            fmt.Println("Error, please try again")
        }
    }
}