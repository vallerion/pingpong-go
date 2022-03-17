package entities

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
	"image/color"
)

type Button struct {
	width, height, borderSize, fontHeight int
	text                                  string
	color                                 color.Color
	borders                               []*ImageAndOptions
	contentImage                          *ImageAndOptions
	font                                  font.Face
	rect                                  *image.Rectangle
	mouseDown                             bool
	clickHandler                          func()
}

func CreateButton(text string, x, y, width, height, borderSize int, fontType *opentype.Font) *Button {
	borders := make([]*ImageAndOptions, 0)
	textColor := color.RGBA{R: 189, G: 195, B: 199, A: 255}

	topBorder := ebiten.NewImage(width, borderSize)
	topBorder.Fill(textColor)
	topBorderOpts := &ebiten.DrawImageOptions{}
	topBorderOpts.GeoM.Translate(float64(x), float64(y))

	bottomBorder := ebiten.NewImage(width, borderSize)
	bottomBorder.Fill(textColor)
	bottomBorderOpts := &ebiten.DrawImageOptions{}
	bottomBorderOpts.GeoM.Translate(float64(x), float64(y+height-borderSize))

	leftBorder := ebiten.NewImage(borderSize, height-borderSize*2)
	leftBorder.Fill(textColor)
	leftBorderOpts := &ebiten.DrawImageOptions{}
	leftBorderOpts.GeoM.Translate(float64(x), float64(y+borderSize))

	rightBorder := ebiten.NewImage(borderSize, height-borderSize*2)
	rightBorder.Fill(textColor)
	rightBorderOpts := &ebiten.DrawImageOptions{}
	rightBorderOpts.GeoM.Translate(float64(x+width-borderSize), float64(y+borderSize))

	borders = append(borders, &ImageAndOptions{topBorder, topBorderOpts})
	borders = append(borders, &ImageAndOptions{bottomBorder, bottomBorderOpts})
	borders = append(borders, &ImageAndOptions{leftBorder, leftBorderOpts})
	borders = append(borders, &ImageAndOptions{rightBorder, rightBorderOpts})

	contentImage := ebiten.NewImage(width-borderSize*2, height-borderSize*2)
	contentImageOpts := &ebiten.DrawImageOptions{}
	contentImageOpts.GeoM.Translate(float64(x+borderSize), float64(y+borderSize))

	faceFont, _ := opentype.NewFace(fontType, &opentype.FaceOptions{
		Size:    15,
		DPI:     72,
		Hinting: font.HintingNone,
	})

	rect := image.Rect(x, y, x+width, y+height)

	b, _, _ := faceFont.GlyphBounds('O')
	fontHeight := (b.Max.Y - b.Min.Y).Ceil()

	return &Button{
		width,
		height,
		borderSize,
		fontHeight,
		text,
		textColor,
		borders,
		&ImageAndOptions{contentImage, contentImageOpts},
		faceFont,
		&rect,
		false,
		func() {},
	}
}

func (b *Button) Update() error {
	pt := image.Pt(ebiten.CursorPosition())
	if pt.In(*b.rect) {
		b.color = color.White
		//ebiten.SetCursorShape(ebiten.CursorShapePointer)

		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			b.mouseDown = true
		} else {
			if b.mouseDown {
				b.clickHandler()
			}
			b.mouseDown = false
		}
	} else {
		b.mouseDown = false
		b.color = color.RGBA{R: 189, G: 195, B: 199, A: 255}
		//ebiten.SetCursorShape(ebiten.CursorShapeDefault)
	}

	return nil
}

func (b *Button) Draw(image *ebiten.Image) {
	for _, border := range b.borders {
		border.Image.Fill(b.color)
		image.DrawImage(border.Image, border.Options)
	}

	image.DrawImage(b.contentImage.Image, b.contentImage.Options)

	bounds, _ := font.BoundString(b.font, b.text)
	w := (bounds.Max.X - bounds.Min.X).Ceil()

	text.Draw(
		b.contentImage.Image,
		b.text,
		b.font,
		b.width/2-w/2,
		b.height/2,
		b.color,
	)
}

func (b *Button) OnClick(clickHandler func()) {
	b.clickHandler = clickHandler
}
