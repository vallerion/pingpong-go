package screens

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/vallerion/pingpong-go/consts"
	"github.com/vallerion/pingpong-go/entities"
	"github.com/vallerion/pingpong-go/resources"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"math/rand"
	"time"
)

type GameScreen struct {
	leftPlayer                                              *entities.Player
	rightPlayer                                             *entities.Player
	balls                                                   []*entities.Ball
	border                                                  *entities.Border
	font                                                    font.Face
	themeAudioPlayer, goalAudioPlayer, collisionAudioPlayer *audio.Player
	pause, keySpacePressed                                  bool
}

func CreateGameScreen(fontType *opentype.Font, audioContext *audio.Context) *GameScreen {
	faceFont, _ := opentype.NewFace(fontType, &opentype.FaceOptions{
		Size:    26,
		DPI:     72,
		Hinting: font.HintingNone,
	})

	gameAudio, _ := mp3.DecodeWithSampleRate(consts.SampleRate, resources.Resources.Get("game"))
	s := audio.NewInfiniteLoopWithIntro(gameAudio, consts.SampleRate*170, consts.SampleRate*170)
	themeAudioPlayer, _ := audioContext.NewPlayer(s)
	themeAudioPlayer.SetVolume(0.4)

	goalAudio, _ := wav.DecodeWithSampleRate(consts.SampleRate, resources.Resources.Get("goal"))
	goalAudioPlayer, _ := audioContext.NewPlayer(goalAudio)
	goalAudioPlayer.SetVolume(1)

	collisionAudio, _ := wav.DecodeWithSampleRate(consts.SampleRate, resources.Resources.Get("collision"))
	collisionAudioPlayer, _ := audioContext.NewPlayer(collisionAudio)
	collisionAudioPlayer.SetVolume(0.75)

	gs := &GameScreen{
		entities.CreatePlayer(consts.GameZoneLeft+50, consts.GameZoneVerticalCenter-consts.PlayerHeight/2),
		entities.CreatePlayer(consts.GameZoneRight-(50+consts.PlayerWidth), consts.GameZoneVerticalCenter-consts.PlayerHeight/2),
		make([]*entities.Ball, 0),
		entities.CreateBorder(),
		faceFont,
		themeAudioPlayer,
		goalAudioPlayer,
		collisionAudioPlayer,
		false,
		false,
	}
	gs.addBall()

	return gs
}

func (s *GameScreen) addBall() {
	s.balls = append(s.balls, entities.CreateBall())
}

func (s *GameScreen) Start() {
	s.themeAudioPlayer.Rewind()
	s.themeAudioPlayer.Play()
}

func (s *GameScreen) End() {
	s.themeAudioPlayer.Close()
}

func (s *GameScreen) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		s.keySpacePressed = true
		//s.pause = !s.pause
	} else if s.keySpacePressed {
		s.pause = !s.pause
		s.keySpacePressed = false
	}

	if s.pause {
		return nil
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

			s.goalAudioPlayer.Rewind()
			s.goalAudioPlayer.Play()
			s.themeAudioPlayer.SetVolume(0.1)
			timer := time.NewTimer(time.Second)
			go func() {
				<-timer.C
				s.themeAudioPlayer.SetVolume(0.4)
			}()
		}

		if ball.RightGoal() {
			ball.ResetPosition()
			s.leftPlayer.Score++

			if len(s.balls) > 1 {
				ball.IsDisplay = false
			}

			s.goalAudioPlayer.Rewind()
			s.goalAudioPlayer.Play()
			s.themeAudioPlayer.SetVolume(0.1)
			timer := time.NewTimer(6 * time.Second)
			go func() {
				<-timer.C
				s.themeAudioPlayer.SetVolume(0.4)
			}()
		}

		ball.Update()

		if s.leftPlayer.GetRect().Overlaps(*ball.GetRect()) {
			ball.Repel(s.leftPlayer.GetRect())
			s.collisionAudioPlayer.Rewind()
			s.collisionAudioPlayer.Play()
		}

		if s.rightPlayer.GetRect().Overlaps(*ball.GetRect()) {
			ball.Repel(s.rightPlayer.GetRect())
			s.collisionAudioPlayer.Rewind()
			s.collisionAudioPlayer.Play()
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
	text.Draw(image, scoreText, s.font, consts.GameZoneHorizontalCenter-20, 35, color.White)

	s.border.Draw(image)
	s.leftPlayer.Draw(image)
	s.rightPlayer.Draw(image)
	for _, ball := range s.balls {
		ball.Draw(image)
	}

	if s.pause {
		s.drawPause(image)
	}
}

func (s *GameScreen) drawPause(image *ebiten.Image) {
	t := "Pause"
	bounds, _ := font.BoundString(s.font, t)
	w := (bounds.Max.X - bounds.Min.X).Ceil()

	text.Draw(image, t, s.font, consts.GameZoneHorizontalCenter-w/2, consts.GameZoneVerticalCenter, color.White)
}
