package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
	"math/rand"
)

const (
	screenHeight             = 768
	screenWidth              = 1366
	gameFieldGap             = 20
	gameZoneLeft             = gameFieldGap
	gameZoneRight            = screenWidth - gameFieldGap
	gameZoneTop              = gameFieldGap + 50
	gameZoneBottom           = screenHeight - gameFieldGap
	gameZoneVerticalCenter   = (gameZoneTop + gameZoneBottom) / 2
	gameZoneHorizontalCenter = (gameZoneRight + gameZoneLeft) / 2
)

var (
	leftPlayer  *Player
	rightPlayer *Player
	balls       []*Ball
	border      *Border
	mainFont    font.Face
)

func init() {
	leftPlayer = CreatePlayer(gameZoneLeft+50, gameZoneVerticalCenter-playerHeight/2)
	rightPlayer = CreatePlayer(gameZoneRight-(50+playerWidth), gameZoneVerticalCenter-playerHeight/2)
	border = CreateBorder()

	balls = make([]*Ball, 0)
	addBall()

	fontObj, _ := opentype.Parse(goregular.TTF)
	mainFont, _ = opentype.NewFace(fontObj, &opentype.FaceOptions{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingNone,
	})
}

func addBall() {
	balls = append(balls, CreateBall())
}

type Game struct{}

func (g *Game) Update() error {
	go func() {
		if rand.Intn(1000) == 50 {
			addBall()
		}
	}()

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		leftPlayer.MoveUp()
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		leftPlayer.MoveDown()
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		rightPlayer.MoveUp()
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		rightPlayer.MoveDown()
	}

	tempBalls := make([]*Ball, 0)
	for _, ball := range balls {
		if ball.IsDisplay == false {
			continue
		}

		if ball.LeftGoal() {
			ball.ResetPosition()
			rightPlayer.Score++

			if len(balls) > 1 {
				ball.IsDisplay = false
			}
		}

		if ball.RightGoal() {
			ball.ResetPosition()
			leftPlayer.Score++

			if len(balls) > 1 {
				ball.IsDisplay = false
			}
		}

		ball.Update()

		if leftPlayer.GetRect().Overlaps(*ball.GetRect()) {
			ball.Repel(leftPlayer.GetRect())
		}

		if rightPlayer.GetRect().Overlaps(*ball.GetRect()) {
			ball.Repel(rightPlayer.GetRect())
		}

		for _, subBall := range balls {
			if ball == subBall {
				continue
			}

			if subBall.GetRect().Overlaps(*ball.GetRect()) {
				ball.Repel(subBall.GetRect())
			}
		}

		if ball.IsDisplay {
			tempBalls = append(tempBalls, ball)
		}
	}

	balls = tempBalls

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	msg := fmt.Sprintf("TPS: %0.2f FPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)

	scoreText := fmt.Sprintf("%d:%d", leftPlayer.Score, rightPlayer.Score)
	text.Draw(screen, scoreText, mainFont, gameZoneHorizontalCenter-20, 40, color.White)

	border.Draw(screen)
	leftPlayer.Draw(screen)
	rightPlayer.Draw(screen)
	for _, ball := range balls {
		ball.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
