package service

import (
	"image"
	"image/color"
	"io/ioutil"
	"math/rand"
	"time"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

const (
	Width      = 200
	Height     = 80
	FontFile   = "arial.ttf"
	FontSize   = 40
	NumChars   = 6
	Characters = "0123456789012345678901234567890123456789z0123456789"
)

func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := range result {
		result[i] = Characters[rand.Intn(len(Characters))]
	}
	return string(result)
}

func CreateImage(text string) (image.Image, error) {
	fontBytes, err := ioutil.ReadFile(FontFile)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}

	rgba := image.NewRGBA(image.Rect(0, 0, Width, Height))
	drawContext := freetype.NewContext()
	drawContext.SetDPI(72)
	drawContext.SetFont(font)
	drawContext.SetFontSize(FontSize)
	drawContext.SetClip(rgba.Bounds())
	drawContext.SetDst(rgba)
	drawContext.SetSrc(image.NewUniform(color.Black))

	pt := freetype.Pt(10, 50)
	_, err = drawContext.DrawString(text, pt)
	if err != nil {
		return nil, err
	}

	return rgba, nil
}
