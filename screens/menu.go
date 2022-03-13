package screens

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/vallerion/pingpong-go/consts"
	"github.com/vallerion/pingpong-go/entities"
	"golang.org/x/image/font/opentype"
)

type Menu struct {
	buttons []*entities.Button
}

func CreateMenuScreen(font *opentype.Font) *Menu {
	buttons := make([]*entities.Button, 0)

	buttons = append(
		buttons,
		entities.CreateButton(
			"Offline game",
			consts.ScreenWidth/2-100,
			consts.ScreenHeight/4,
			200,
			75,
			5,
			font,
		),
	)

	return &Menu{buttons}
}

func (m *Menu) Update() error {
	for _, button := range m.buttons {
		if err := button.Update(); err != nil {
			return err
		}
	}

	return nil
}

func (m *Menu) Draw(image *ebiten.Image) {
	for _, button := range m.buttons {
		button.Draw(image)
	}
}
