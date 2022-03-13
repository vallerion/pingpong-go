package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/vallerion/pingpong-go/consts"
	"image/color"
)

type Border struct {
	//image *ebiten.Image
}

func CreateBorder() *Border {
	return &Border{}
}

func (p *Border) Draw(screen *ebiten.Image) {
	topBorder := ebiten.NewImage(consts.ScreenWidth, consts.GameFieldGap)
	topBorder.Fill(color.White)
	topBorderOpts := &ebiten.DrawImageOptions{}
	topBorderOpts.GeoM.Translate(float64(0), float64(consts.GameZoneTop-consts.GameFieldGap))

	bottomBorder := ebiten.NewImage(consts.ScreenWidth, consts.GameFieldGap)
	bottomBorder.Fill(color.White)
	bottomBorderOpts := &ebiten.DrawImageOptions{}
	bottomBorderOpts.GeoM.Translate(float64(0), float64(consts.GameZoneBottom))

	leftTopBorder := ebiten.NewImage(consts.GameFieldGap, consts.SideBorderHeight)
	leftTopBorder.Fill(color.White)
	leftTopBorderOpts := &ebiten.DrawImageOptions{}
	leftTopBorderOpts.GeoM.Translate(float64(0), float64(consts.GameZoneTop))

	leftBottomBorder := ebiten.NewImage(consts.GameFieldGap, consts.SideBorderHeight)
	leftBottomBorder.Fill(color.White)
	leftBottomBorderOpts := &ebiten.DrawImageOptions{}
	leftBottomBorderOpts.GeoM.Translate(float64(0), float64(consts.GateBottom))

	rightTopBorder := ebiten.NewImage(consts.GameFieldGap, consts.SideBorderHeight)
	rightTopBorder.Fill(color.White)
	rightTopBorderOpts := &ebiten.DrawImageOptions{}
	rightTopBorderOpts.GeoM.Translate(float64(consts.ScreenWidth-consts.GameFieldGap), float64(consts.GameZoneTop))

	rightBottomBorder := ebiten.NewImage(consts.GameFieldGap, consts.SideBorderHeight)
	rightBottomBorder.Fill(color.White)
	rightBottomBorderOpts := &ebiten.DrawImageOptions{}
	rightBottomBorderOpts.GeoM.Translate(float64(consts.ScreenWidth-consts.GameFieldGap), float64(consts.GateBottom))

	screen.DrawImage(topBorder, topBorderOpts)
	screen.DrawImage(bottomBorder, bottomBorderOpts)
	screen.DrawImage(leftTopBorder, leftTopBorderOpts)
	screen.DrawImage(leftBottomBorder, leftBottomBorderOpts)
	screen.DrawImage(rightTopBorder, rightTopBorderOpts)
	screen.DrawImage(rightBottomBorder, rightBottomBorderOpts)
}
