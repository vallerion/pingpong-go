package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/vallerion/pingpong-go/consts"
	"github.com/vallerion/pingpong-go/entities"
	"github.com/vallerion/pingpong-go/screens"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
	"log"
)

var (
	//gameScreen *screens.GameScreen
	menuScreen *screens.Menu
	fontObj    *opentype.Font
)

func init() {
	fontObj, _ = opentype.Parse(goregular.TTF)

	//gameScreen = screens.CreateScreen(mainFont)
	//menuScreen = screens.CreateMenuScreen(fontObj)
}

func createGame() *Game {
	m := screens.CreateMenuScreen(fontObj)

	g := &Game{m}

	m.OnStartOffline(func() {
		g.currentScreen = screens.CreateGameScreen(fontObj)
	})

	return g
}

type Game struct {
	currentScreen entities.Entity
}

func (g *Game) Update() error {
	_, isGame := g.currentScreen.(*screens.GameScreen)
	if ebiten.IsKeyPressed(ebiten.KeyEscape) && isGame {
		g.currentScreen = screens.CreateMenuScreen(fontObj)
	}

	return g.currentScreen.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentScreen.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return consts.ScreenWidth, consts.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(consts.ScreenWidth, consts.ScreenHeight)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(createGame()); err != nil {
		log.Fatal(err)
	}
}
