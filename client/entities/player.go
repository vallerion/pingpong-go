package entities

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/vallerion/pingpong-go/client/consts"
	"image"
	"image/color"
)

type Player struct {
	image *ebiten.Image
	rect  *image.Rectangle
	Score uint
}

func CreatePlayer(x, y float64) *Player {
	rect := image.Rect(int(x), int(y), int(x+consts.PlayerWidth), int(y+consts.PlayerHeight))
	return &Player{ebiten.NewImage(consts.PlayerWidth, consts.PlayerHeight), &rect, 0}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.image.Fill(color.White)
	leftOpts := &ebiten.DrawImageOptions{}
	leftOpts.GeoM.Translate(float64(p.rect.Min.X), float64(p.rect.Min.Y))
	screen.DrawImage(p.image, leftOpts)
}

func (p *Player) MoveUp() {
	if p.rect.Min.Y <= consts.GameZoneTop {
		return
	}
	p.rect.Min.Y -= 5
	p.rect.Max.Y -= 5
}

func (p *Player) MoveDown() {
	if p.rect.Max.Y >= consts.GameZoneBottom {
		return
	}
	p.rect.Min.Y += 5
	p.rect.Max.Y += 5
}

func (p *Player) GetRect() *image.Rectangle {
	return p.rect
}
