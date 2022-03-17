package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/vallerion/pingpong-go/client/consts"
	"github.com/vallerion/pingpong-go/client/screens"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
	"log"
)

var (
	fontObj      *opentype.Font
	audioContext *audio.Context
)

const (
	screenMenu = iota
	screenGameOffline
)

func init() {
	fontObj, _ = opentype.Parse(goregular.TTF)
	audioContext = audio.NewContext(consts.SampleRate)
}

func createGame() *Game {
	m := screens.CreateMenuScreen(fontObj, audioContext)

	scr := make(map[int]screens.Screen, 2)
	scr[screenMenu] = m
	scr[screenGameOffline] = screens.CreateGameScreen(fontObj, audioContext)

	g := &Game{screenMenu, scr}

	m.OnStartOffline(func() {
		g.screens[g.currentScreen].End()
		g.currentScreen = screenGameOffline
		g.screens[g.currentScreen].Start()
	})
	m.Start()

	return g
}

type Game struct {
	currentScreen int
	screens       map[int]screens.Screen
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) && g.currentScreen == screenGameOffline {
		g.screens[g.currentScreen].End()
		g.currentScreen = screenMenu
		g.screens[g.currentScreen].Start()
	}

	return g.screens[g.currentScreen].Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	msg := fmt.Sprintf("TPS: %0.2f FPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)

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
