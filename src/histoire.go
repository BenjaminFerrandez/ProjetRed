package main

import (
	"fmt"

	"image/color"
	"io/ioutil"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font/opentype"
)
type Button struct {
    X, Y, Width, Height int
    Label              string
    Image              *ebiten.Image
    Active             bool
}
var (
    historyLines = []string{
        "\n                        Welcome in our world !\n",
        "           A world where people have some magical power.\n",
        "            You will be one of them to complete a mission.\n",
        "The mission is to kill a monster but we know one specific thing about him.\n",
        "     The thing that you need to know is that when you kill him,",
        "              he revives with a little bit more of life.\n",
        "   The only way to kill him definitely is to kill him 4 times in a row.\n",
        "To succeed, you can buy some items that will help you, so do not hesitate",
        "              (if you have enough money, of course)\n",
        "                        Good luck and enjoy !",
    }
    currentLineIndex int
    displayedText    string
    lastUpdateTime   time.Time
    gameState         int
    GameStateHistory  = 0
    GameStateStart    = 1
    buttons []Button
)

func init() {
    // Download image
    img, _, err := ebitenutil.NewImageFromFile("/home/alexandre/ProjetRed (copy graphique)/src/screen/edit1024x768.png", ebiten.FilterDefault) // Укажите путь к вашему изображению фона
    if err != nil {
        panic(err)
    }
    backgroundImg = img


    fontData, err := ioutil.ReadFile("/home/alexandre/ProjetRed (copy graphique)/src/screen/MenuCard.ttf")
    if err != nil {
        log.Fatal(err)
    }

    font, err := opentype.Parse(fontData)
    if err != nil {
        log.Fatal(err)
    }

    const dpi = 72
    Font, err = opentype.NewFace(font, &opentype.FaceOptions{
        Size:    14, 
        DPI:     dpi,
       
    })
    if err != nil {
        log.Fatal(err)
    }
    // Load images for buttons
playImg, _, err := ebitenutil.NewImageFromFile("/home/alexandre/ProjetRed (copy graphique)/src/screen/23-04-47-748.png", ebiten.FilterDefault)
if err != nil {
    log.Fatal(err)
}

returnImg, _, err := ebitenutil.NewImageFromFile("/home/alexandre/ProjetRed (copy graphique)/src/screen/23-04-47-748.png", ebiten.FilterDefault)
if err != nil {
    log.Fatal(err)
}

quitImg, _, err := ebitenutil.NewImageFromFile("/home/alexandre/ProjetRed (copy graphique)/src/screen/23-04-47-748.png", ebiten.FilterDefault)
if err != nil {
    log.Fatal(err)
}

buttons = []Button{
    {X: 400, Y: 450, Width: 200, Height: 40, Label: "Play", Image: playImg, Active: true},
    {X: 400, Y: 500, Width: 300, Height: 40, Label: "Return", Image: returnImg, Active: true},
    {X: 400, Y: 550, Width: 200, Height: 40, Label: "Quit", Image: quitImg, Active: true},
}
}


func history(screen *ebiten.Image) error {

    screen.Fill(color.Black)
    if !gameStarted {
        if err := screen.DrawImage(backgroundImg, nil); err != nil {
            return err
        }
        text.Draw(screen, "D E A T H     B A C K", Font, 350, 60, color.White)
        if currentLineIndex < len(historyLines) {
            if time.Since(lastUpdateTime) >= time.Second*2 {
                displayedText += historyLines[currentLineIndex] + "\n"
                currentLineIndex++
                lastUpdateTime = time.Now()
            }
        }
    
     
        textX := 100 
        textY := 80 
    
        text.Draw(screen, displayedText, Font, textX, textY, color.White)

       for _, button := range buttons {
        color := color.White
        if button.Active {
            if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
            mouseX, mouseY := ebiten.CursorPosition()
            if mouseX >= button.X && mouseX <= button.X+button.Width && mouseY >= button.Y && mouseY <= button.Y+button.Height {
                // button click 
                if button.Label == "Play" {
                    return fmt.Errorf("Play")
                } else if button.Label == "Return" {
                  
                    gameState = GameStateHistory
                    displayedText = ""
                    currentLineIndex = 0
                } else if button.Label == "Quit" {
                    fmt.Println("Goodbye!")
                    return fmt.Errorf("Quit")
                }
            }
        }
    }

    // Draw the button image
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(float64(button.X), float64(button.Y))
    screen.DrawImage(button.Image, op)

    // Draw the button label
    text.Draw(screen, button.Label, Font, button.X+45, button.Y+35, color)
}
    }
    
    // Здесь можно обновлять состояние игры и рисовать графику
    // ...
    
    return nil
}
