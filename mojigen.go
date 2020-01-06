package mojigen

import (
	"image"
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// Mojigen は文字列を受け取って、それを印字したpng画像のバイト列を生成します
func Mojigen(ss []string, c Config) (*image.RGBA, error) {
	ttfbytes, err := ioutil.ReadFile(c.FontPath)
	if err != nil {
		return nil, err
	}
	ttf, err := truetype.Parse(ttfbytes)
	if err != nil {
		return nil, err
	}

	opt := truetype.Options{
		Size: c.FontSize,
	}
	face := truetype.NewFace(ttf, &opt)

	dst := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))
	draw.Draw(dst, dst.Bounds(), image.NewUniform(c.ColorBackground), image.Point{}, draw.Src)
	d := &font.Drawer{
		Dst:  dst,
		Src:  image.NewUniform(c.ColorFront),
		Face: face,
	}

	for i, s := range ss {
		d.Dot.X = fixed.I(c.MarginLeft)
		d.Dot.Y = fixed.I(c.Height/2 - c.LineHeight*len(ss)/2 + c.LineHeight*(i+1) - c.LineHeight/5)
		d.DrawString(s)
	}

	return dst, err
}
