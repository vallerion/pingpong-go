package screens

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/vallerion/pingpong-go/consts"
	"github.com/vallerion/pingpong-go/entities"
	"github.com/vallerion/pingpong-go/resources"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
)

type Menu struct {
	onStartOnlineBtn  *entities.Button
	onStartOfflineBtn *entities.Button
	onExitBtn         *entities.Button
	audioPlayer       *audio.Player
	fontFace          font.Face
}

func CreateMenuScreen(fontType *opentype.Font, audioContext *audio.Context) *Menu {
	x := float64(consts.ScreenWidth)*0.3 - 100

	offline := entities.CreateButton(
		"Offline game",
		int(x),
		consts.ScreenHeight/4,
		200,
		75,
		5,
		fontType,
	)

	online := entities.CreateButton(
		"Multiplayer",
		int(x),
		consts.ScreenHeight/4+100,
		200,
		75,
		5,
		fontType,
	)

	exit := entities.CreateButton(
		"Exit",
		int(x),
		consts.ScreenHeight/4+200,
		200,
		75,
		5,
		fontType,
	)

	d, _ := mp3.DecodeWithSampleRate(consts.SampleRate, resources.Resources.Get("menu"))
	//s := audio.NewInfiniteLoopWithIntro(d, consts.SampleRate, consts.SampleRate)
	audioPlayer, _ := audioContext.NewPlayer(d)

	fontFace, _ := opentype.NewFace(fontType, &opentype.FaceOptions{
		Size:    28,
		DPI:     72,
		Hinting: font.HintingNone,
	})

	return &Menu{online, offline, exit, audioPlayer, fontFace}
}

func (m *Menu) Update() error {
	m.onStartOnlineBtn.Update()
	m.onStartOfflineBtn.Update()
	//m.onExitBtn.Update()

	return nil
}

func (m *Menu) Draw(image *ebiten.Image) {
	m.onStartOnlineBtn.Draw(image)
	m.onStartOfflineBtn.Draw(image)
	//m.onExitBtn.Draw(image)

	m.drawControlInfo(image)
}

func (m *Menu) drawControlInfo(image *ebiten.Image) {
	x := float64(consts.ScreenWidth) * 0.5

	controlText := "How to play:\n" +
		"W,A,S,D to move left player\n" +
		"Arrows to move right player\n" +
		"Space to pause"
	text.Draw(image, controlText, m.fontFace, int(x), consts.ScreenHeight/4+20, color.White)
}

func (m *Menu) Start() {
	m.audioPlayer.Rewind()
	m.audioPlayer.Play()
}

func (m *Menu) End() {
	m.audioPlayer.Close()
}

func (m *Menu) OnStartOnline(handler func()) {
	m.onStartOnlineBtn.OnClick(handler)
}

func (m *Menu) OnStartOffline(handler func()) {
	m.onStartOfflineBtn.OnClick(handler)
}

func (m *Menu) OnExit(handler func()) {
	m.onExitBtn.OnClick(handler)
}
