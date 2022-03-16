package screens

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Screen interface {
	Update() error
	Draw(*ebiten.Image)
	Start()
	End()
}
