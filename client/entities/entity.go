package entities

import "github.com/hajimehoshi/ebiten"

type Entity interface {
	Update() error
	Draw(*ebiten.Image)
}
