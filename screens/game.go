package screens

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/vallerion/pingpong-go/consts"
	"github.com/vallerion/pingpong-go/entities"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"math/rand"
)

type GameScreen struct {
	leftPlayer  *entities.Player
	rightPlayer *entities.Player
	balls       []*entities.Ball
	border      *entities.Border
	font        font.Face
	pause       bool
}

func CreateGameScreen(fontType *opentype.Font) *GameScreen {
	faceFont, _ := opentype.NewFace(fontType, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingNone,
	})

	image := &GameScreen{
		entities.CreatePlayer(consts.GameZoneLeft+50, consts.GameZoneVerticalCenter-consts.PlayerHeight/2),
		entities.CreatePlayer(consts.GameZoneRight-(50+consts.PlayerWidth), consts.GameZoneVerticalCenter-consts.PlayerHeight/2),
		make([]*entities.Ball, 0),
		entities.CreateBorder(),
		faceFont,
		false,
	}

	image.addBall()
	return image
}

func (s *GameScreen) addBall() {
	s.balls = append(s.balls, entities.CreateBall())
}

func (s *GameScreen) Update() error {
	if s.pause {
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		s.pause = !s.pause
	}

	go func() {
		if rand.Intn(3000) == 50 {
			s.addBall()
		}
	}()
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		s.leftPlayer.MoveUp()
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		s.leftPlayer.MoveDown()
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		s.rightPlayer.MoveUp()
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		s.rightPlayer.MoveDown()
	}

	tempBalls := make([]*entities.Ball, 0)
	for _, ball := range s.balls {
		if ball.IsDisplay == false {
			continue
		}

		if ball.LeftGoal() {
			ball.ResetPosition()
			s.rightPlayer.Score++

			if len(s.balls) > 1 {
				ball.IsDisplay = false
			}
		}

		if ball.RightGoal() {
			ball.ResetPosition()
			s.leftPlayer.Score++

			if len(s.balls) > 1 {
				ball.IsDisplay = false
			}
		}

		ball.Update()

		if s.leftPlayer.GetRect().Overlaps(*ball.GetRect()) {
			ball.Repel(s.leftPlayer.GetRect())
		}

		if s.rightPlayer.GetRect().Overlaps(*ball.GetRect()) {
			ball.Repel(s.rightPlayer.GetRect())
		}

		for _, subBall := range s.balls {
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

	// remove those that are not showing
	s.balls = tempBalls

	return nil
}

func (s *GameScreen) Draw(image *ebiten.Image) {
	scoreText := fmt.Sprintf("%d:%d", s.leftPlayer.Score, s.rightPlayer.Score)
	text.Draw(image, scoreText, s.font, consts.GameZoneHorizontalCenter-20, 40, color.White)

	s.border.Draw(image)
	s.leftPlayer.Draw(image)
	s.rightPlayer.Draw(image)
	for _, ball := range s.balls {
		ball.Draw(image)
	}
}
