package main

import (
	"image/png"
	"log"
	"os"

	"github.com/aimof/mojigen"
)

func main() {
	if len(os.Args) <= 2 {
		log.Println("第三引数以降で文面を指定してください")
		os.Exit(1)
	}
	ss := os.Args[2:]
	img, err := mojigen.Mojigen(ss, mojigen.DefaultConfig)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	out, err := os.Create(os.Args[1])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer out.Close()

	if err := png.Encode(out, img); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
