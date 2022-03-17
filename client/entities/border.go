package entities

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/vallerion/pingpong-go/client/consts"
	"image/color"
)

type Border struct {
	borders []*ImageAndOptions
}

func CreateBorder() *Border {
	borders := make([]*ImageAndOptions, 0)

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

	borders = append(borders, &ImageAndOptions{topBorder, topBorderOpts})
	borders = append(borders, &ImageAndOptions{bottomBorder, bottomBorderOpts})
	borders = append(borders, &ImageAndOptions{leftTopBorder, leftTopBorderOpts})
	borders = append(borders, &ImageAndOptions{leftBottomBorder, leftBottomBorderOpts})
	borders = append(borders, &ImageAndOptions{rightTopBorder, rightTopBorderOpts})
	borders = append(borders, &ImageAndOptions{rightBottomBorder, rightBottomBorderOpts})

	return &Border{borders}
}

func (b *Border) Draw(screen *ebiten.Image) {
	for _, border := range b.borders {
		screen.DrawImage(border.Image, border.Options)
	}
}
