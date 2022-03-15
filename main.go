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
	fontObj *opentype.Font
)

const (
	screenMenu = iota
	screenGameOffline
)

func init() {
	fontObj, _ = opentype.Parse(goregular.TTF)
}

func createGame() *Game {
	m := screens.CreateMenuScreen(fontObj)

	scr := make(map[int]entities.Entity, 2)
	scr[screenMenu] = m
	scr[screenGameOffline] = screens.CreateGameScreen(fontObj)

	g := &Game{screenMenu, scr}

	m.OnStartOffline(func() {
		g.currentScreen = screenGameOffline
	})

	return g
}

type Game struct {
	currentScreen int
	screens       map[int]entities.Entity
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) && g.currentScreen == screenGameOffline {
		g.currentScreen = screenMenu
	}

	return g.screens[g.currentScreen].Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.screens[g.currentScreen].Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return consts.ScreenWidth, consts.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(consts.ScreenWidth, consts.ScreenHeight)
	ebiten.SetWindowTitle("Ping Pong Online")

	if err := ebiten.RunGame(createGame()); err != nil {
		log.Fatal(err)
	}
}
