package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
    historyLines = []string{
        "Welcome in our world !",
        "A world where people have some magical power.",
        "You will be one of them to complete a mission",
        "The mission is to kill a monster but we know one specific thing about him.",
        "The thing that you need to know is that when you kill him, he revives with a little bit more of life.",
        "The only way to kill him definitely is to kill him 4 times in a row",
        "To succeed, you can buy some items that will help you, so do not hesitate (if you have enough money, of course)",
        "Good luck and enjoy !",
    }
    currentLineIndex int
    displayedText    string
    lastUpdateTime   time.Time
)

func history(screen *ebiten.Image) error {

    screen.Fill(color.Black)


    if currentLineIndex < len(historyLines) {
        if time.Since(lastUpdateTime) >= time.Second*2 {
            displayedText += historyLines[currentLineIndex] + "\n"
            currentLineIndex++
            lastUpdateTime = time.Now()
        }
    }

 
    textX := 100 
    textY := 100 


    ebitenutil.DebugPrintAt(screen, displayedText, textX, textY)

    return nil
}