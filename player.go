package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

const (
	playerWidth  = 16
	playerHeight = 48
)

type Player struct {
	image *ebiten.Image
	rect  *image.Rectangle
	Score uint
}

func CreatePlayer(x, y float64) *Player {
	rect := image.Rect(int(x), int(y), int(x+playerWidth), int(y+playerHeight))
	return &Player{ebiten.NewImage(playerWidth, playerHeight), &rect, 0}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.image.Fill(color.White)
	leftOpts := &ebiten.DrawImageOptions{}
	leftOpts.GeoM.Translate(float64(p.rect.Min.X), float64(p.rect.Min.Y))
	screen.DrawImage(p.image, leftOpts)
}

func (p *Player) MoveUp() {
	if p.rect.Min.Y <= gameZoneTop {
		return
	}
	p.rect.Min.Y -= 5
	p.rect.Max.Y -= 5
}

func (p *Player) MoveDown() {
	if p.rect.Max.Y >= gameZoneBottom {
		return
	}
	p.rect.Min.Y += 5
	p.rect.Max.Y += 5
}

func (p *Player) GetRect() *image.Rectangle {
	return p.rect
}
