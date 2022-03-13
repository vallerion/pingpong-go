package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/vallerion/pingpong-go/consts"
	"github.com/vallerion/pingpong-go/screens"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
	"log"
)

var (
	//gameScreen *screens.GameScreen
	menuScreen *screens.Menu
	mainFont   font.Face
)

func init() {
	fontObj, _ := opentype.Parse(goregular.TTF)
	mainFont, _ = opentype.NewFace(fontObj, &opentype.FaceOptions{
		Size:    consts.TextSize,
		DPI:     72,
		Hinting: font.HintingNone,
	})

	//gameScreen = screens.CreateScreen(mainFont)
	menuScreen = screens.CreateMenuScreen(fontObj)
}

type Game struct{}

func (g *Game) Update() error {
	return menuScreen.Update()
	//return gameScreen.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	menuScreen.Draw(screen)
	//gameScreen.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return consts.ScreenWidth, consts.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(consts.ScreenWidth, consts.ScreenHeight)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
