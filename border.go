package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

var (
	gateHeight       = 120
	sideBorderHeight = (screenHeight-gameZoneTop-gameFieldGap*2)/2 - gateHeight/2
	gateTop          = screenHeight/2 - gateHeight/2 + gameZoneTop
	gateBottom       = screenHeight/2 + gateHeight/2 + gameZoneTop
)

type Border struct {
	//image *ebiten.Image
}

func CreateBorder() *Border {
	return &Border{}
}

func (p *Border) Draw(screen *ebiten.Image) {
	topBorder := ebiten.NewImage(screenWidth, gameFieldGap)
	topBorder.Fill(color.White)
	topBorderOpts := &ebiten.DrawImageOptions{}
	topBorderOpts.GeoM.Translate(float64(0), float64(gameZoneTop-gameFieldGap))

	bottomBorder := ebiten.NewImage(screenWidth, gameFieldGap)
	bottomBorder.Fill(color.White)
	bottomBorderOpts := &ebiten.DrawImageOptions{}
	bottomBorderOpts.GeoM.Translate(float64(0), float64(gameZoneBottom))

	leftTopBorder := ebiten.NewImage(gameFieldGap, sideBorderHeight)
	leftTopBorder.Fill(color.White)
	leftTopBorderOpts := &ebiten.DrawImageOptions{}
	leftTopBorderOpts.GeoM.Translate(float64(0), float64(gameZoneTop))

	leftBottomBorder := ebiten.NewImage(gameFieldGap, sideBorderHeight)
	leftBottomBorder.Fill(color.White)
	leftBottomBorderOpts := &ebiten.DrawImageOptions{}
	leftBottomBorderOpts.GeoM.Translate(float64(0), float64(gateBottom))

	rightTopBorder := ebiten.NewImage(gameFieldGap, sideBorderHeight)
	rightTopBorder.Fill(color.White)
	rightTopBorderOpts := &ebiten.DrawImageOptions{}
	rightTopBorderOpts.GeoM.Translate(float64(screenWidth-gameFieldGap), float64(gameZoneTop))

	rightBottomBorder := ebiten.NewImage(gameFieldGap, sideBorderHeight)
	rightBottomBorder.Fill(color.White)
	rightBottomBorderOpts := &ebiten.DrawImageOptions{}
	rightBottomBorderOpts.GeoM.Translate(float64(screenWidth-gameFieldGap), float64(gateBottom))

	screen.DrawImage(topBorder, topBorderOpts)
	screen.DrawImage(bottomBorder, bottomBorderOpts)
	screen.DrawImage(leftTopBorder, leftTopBorderOpts)
	screen.DrawImage(leftBottomBorder, leftBottomBorderOpts)
	screen.DrawImage(rightTopBorder, rightTopBorderOpts)
	screen.DrawImage(rightBottomBorder, rightBottomBorderOpts)
}
