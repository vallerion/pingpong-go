package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	"math"
	"math/rand"
)

type Ball struct {
	initX, initY, dx, dy float64
	image                *ebiten.Image
	rect                 *image.Rectangle
}

const (
	ballHeight   = 10
	ballWidth    = 10
	ballMaxSpeed = 30
)

func CreateBall(x, y float64) *Ball {
	dx, dy := randDx(), 0.
	rect := image.Rect(int(x), int(y), int(x+ballWidth), int(y+ballHeight))

	return &Ball{x, y, dx, dy, ebiten.NewImage(ballWidth, ballHeight), &rect}
}

func randDx() float64 {
	direction := rand.Intn(2)
	dx := 5 + rand.Intn(3)

	if direction == 1 {
		return float64(-dx)
	}

	return float64(dx)
}

func (p *Ball) ResetPosition() {
	p.dx, p.dy = randDx(), 0
	p.rect.Min.X = int(p.initX)
	p.rect.Min.Y = int(p.initY)
	p.rect.Max.X = int(p.initX + ballWidth)
	p.rect.Max.Y = int(p.initY + ballHeight)
}

func (p *Ball) Draw(screen *ebiten.Image) {
	if math.Abs(p.dx) < 10 {
		p.image.Fill(color.White)
	} else if math.Abs(p.dx) < 20 {
		p.image.Fill(color.RGBA{R: 255, G: 240, B: 0, A: 255})
	} else {
		p.image.Fill(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	}
	leftOpts := &ebiten.DrawImageOptions{}
	leftOpts.GeoM.Translate(float64(p.rect.Min.X), float64(p.rect.Min.Y))
	screen.DrawImage(p.image, leftOpts)
}

func (p *Ball) Update() {
	if p.rect.Max.Y >= gameZoneBottom {
		p.dy = -p.dy
	}
	if p.rect.Min.Y <= gameZoneTop {
		p.dy = -p.dy
	}
	if p.rect.Max.X >= gameZoneRight {
		p.dx = -p.dx
	}
	if p.rect.Min.X <= gameZoneLeft {
		p.dx = -p.dx
	}

	p.rect.Min.X += int(p.dx)
	p.rect.Min.Y += int(p.dy)
	p.rect.Max.X += int(p.dx)
	p.rect.Max.Y += int(p.dy)
}

func (p *Ball) GetRect() *image.Rectangle {
	return p.rect
}

func (p *Ball) Repel(rect *image.Rectangle) {
	horizontalCollision := p.rect.Min.Y >= rect.Min.Y || p.rect.Max.Y <= rect.Max.Y

	rectCenter := (rect.Max.Y + rect.Min.Y) / 2
	ballCenter := (p.rect.Max.Y + p.rect.Min.Y) / 2
	diff := rectCenter - ballCenter

	if horizontalCollision {
		if math.Abs(p.dx) < ballMaxSpeed {
			p.dx = -(p.dx * 1.3)
		} else {
			p.dx = -p.dx
		}
		p.dy = normalizeDiffRepel(float64(diff)) + randomizeDY()
	} else {
		p.dy = -p.dy + randomizeDY()
	}
}

func randomizeDY() float64 {
	return float64(1+rand.Intn(1)) / float64(2+rand.Intn(5))
}

func normalizeDiffRepel(diff float64) float64 {
	constraint := 3.

	if diff < -constraint {
		return constraint
	}
	if diff > constraint {
		return -constraint
	}
	return -diff
}

func (p *Ball) LeftGoal() bool {
	return p.rect.Min.X <= gameZoneLeft && p.rect.Min.Y >= gateTop && p.rect.Max.Y <= gateBottom
}

func (p *Ball) RightGoal() bool {
	return p.rect.Max.X >= gameZoneRight && p.rect.Min.Y >= gateTop && p.rect.Max.Y <= gateBottom
}