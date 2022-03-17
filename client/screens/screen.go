package screens

import (
	"github.com/hajimehoshi/ebiten"
)

type Screen interface {
	Update() error
	Draw(*ebiten.Image)
	Start()
	End()
}
