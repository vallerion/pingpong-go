package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/vallerion/pingpong-go/consts"
	"image"
	"image/color"
	"math"
	"math/rand"
)

type Ball struct {
	initX, initY int
	dx, dy       float64
	image        *ebiten.Image
	rect         *image.Rectangle
	IsDisplay    bool
}

func CreateBall() *Ball {
	x, y := consts.GameZoneHorizontalCenter-consts.BallSize/2, consts.GameZoneVerticalCenter-consts.BallSize/2
	dx, dy := randDx(), 0.
	rect := image.Rect(x, y, x+consts.BallSize, y+consts.BallSize)

	return &Ball{
		x,
		y,
		dx,
		dy,
		ebiten.NewImage(consts.BallSize, consts.BallSize),
		&rect,
		true,
	}
}

func randDx() float64 {
	direction := rand.Intn(2)
	dx := 3 + rand.Intn(3)

	if direction == 1 {
		return float64(-dx)
	}

	return float64(dx)
}

func (p *Ball) ResetPosition() {
	p.dx, p.dy = randDx(), 0
	p.rect.Min.X = p.initX
	p.rect.Min.Y = p.initY
	p.rect.Max.X = int(p.initX) + consts.BallSize
	p.rect.Max.Y = int(p.initY) + consts.BallSize
}

func (p *Ball) Draw(screen *ebiten.Image) {
	if p.IsDisplay == false {
		return
	}

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

	rightImage := ebiten.NewImage(100, 100)
	rightImage.Fill(color.RGBA{R: 255})
	rightImageOpts := &ebiten.DrawImageOptions{}
	rightImageOpts.GeoM.Translate(float64(consts.GameZoneRight), float64(consts.GateTop))
	screen.DrawImage(rightImage, rightImageOpts)

	leftImage := ebiten.NewImage(100, 100)
	leftImage.Fill(color.RGBA{R: 255})
	leftImageOpts := &ebiten.DrawImageOptions{}
	leftImageOpts.GeoM.Translate(float64(consts.GameZoneLeft), float64(consts.GateTop))
	screen.DrawImage(leftImage, leftImageOpts)
}

func (p *Ball) Update() {
	if p.rect.Max.Y >= consts.GameZoneBottom {
		p.dy = -p.dy
	}
	if p.rect.Min.Y <= consts.GameZoneTop {
		p.dy = -p.dy
	}
	if p.rect.Max.X >= consts.GameZoneRight {
		p.dx = -p.dx
	}
	if p.rect.Min.X <= consts.GameZoneLeft {
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
		if math.Abs(p.dx) < consts.BallMaxSpeed {
			p.dx = -(p.dx * 1.15)
		} else {
			p.dx = -p.dx
		}
		p.dy = normalizeDiffRepel(float64(diff)) + randomizeDY()
	} else {
		p.dy = -p.dy + randomizeDY()
	}
}

func randomizeDY() float64 {
	return float64(1+rand.Intn(1)) / float64(2+rand.Intn(3))
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
	return p.rect.Min.X <= consts.GameZoneLeft && p.rect.Min.Y >= consts.GateTop && p.rect.Max.Y <= consts.GateBottom
}

func (p *Ball) RightGoal() bool {
	return p.rect.Max.X >= consts.GameZoneRight && p.rect.Min.Y >= consts.GateTop && p.rect.Max.Y <= consts.GateBottom
}
