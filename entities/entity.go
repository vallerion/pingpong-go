package entities

import "github.com/hajimehoshi/ebiten/v2"

type Entity interface {
	Update() error
	Draw(*ebiten.Image)
}
