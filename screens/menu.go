package screens

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/vallerion/pingpong-go/consts"
	"github.com/vallerion/pingpong-go/entities"
	"golang.org/x/image/font/opentype"
)

type Menu struct {
	onStartOnlineBtn  *entities.Button
	onStartOfflineBtn *entities.Button
	onExitBtn         *entities.Button
}

func CreateMenuScreen(font *opentype.Font) *Menu {
	x := float64(consts.ScreenWidth)*0.3 - 100

	offline := entities.CreateButton(
		"Offline game",
		int(x),
		consts.ScreenHeight/4,
		200,
		75,
		5,
		font,
	)

	online := entities.CreateButton(
		"Multiplayer",
		int(x),
		consts.ScreenHeight/4+100,
		200,
		75,
		5,
		font,
	)

	exit := entities.CreateButton(
		"Exit",
		int(x),
		consts.ScreenHeight/4+200,
		200,
		75,
		5,
		font,
	)

	return &Menu{online, offline, exit}
}

func (m *Menu) Update() error {
	m.onStartOnlineBtn.Update()
	m.onStartOfflineBtn.Update()
	m.onExitBtn.Update()

	return nil
}

func (m *Menu) Draw(image *ebiten.Image) {
	m.onStartOnlineBtn.Draw(image)
	m.onStartOfflineBtn.Draw(image)
	m.onExitBtn.Draw(image)
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
