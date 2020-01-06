package mojigen

import (
	"image/color"
	"os"
)

// Config です。だいたい書いてあるとおりです。
type Config struct {
	FontPath        string
	ColorFront      color.Color
	ColorBackground color.Color
	Width           int
	Height          int
	MarginLeft      int
	LineHeight      int
	FontSize        float64
}

// DefaultConfig もし設定項目がなければこれを使ってください。
var DefaultConfig = Config{
	FontPath:        os.Getenv("GOPATH") + "/src/github.com/aimof/mojigen/Koruri/Koruri-Regular.ttf",
	ColorFront:      color.Black,
	ColorBackground: color.White,
	Width:           506,
	Height:          253,
	MarginLeft:      40,
	LineHeight:      60,
	FontSize:        48,
}
