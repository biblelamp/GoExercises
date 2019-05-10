package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {

	rectangle := "rectangle.png"

	green := color.RGBA{0, 255, 0, 255}

	rectImg := image.NewRGBA(image.Rect(0, 0, 200, 200))

	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

	file, err := os.Create(rectangle)
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	png.Encode(file, rectImg)
}
