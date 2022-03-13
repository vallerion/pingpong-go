package entities

import "github.com/hajimehoshi/ebiten/v2"

type ImageAndOptions struct {
	Image   *ebiten.Image
	Options *ebiten.DrawImageOptions
}
