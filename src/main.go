package main

import "fmt"

func main() {
    for {
        fmt.Println("Bienvenue")
        fmt.Println("Click 's' to start or 'q' to quit")
        var choice string
        fmt.Scanln(&choice)
        switch choice {
        	case "s":
            fmt.Println("Goodluck and enjoy !")
            start()
       		case "q":
            fmt.Println("Goodbye!")
            return
       		default:
            fmt.Println("Error, please try again")
        }
    }
}