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

	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}
	board := image.NewRGBA(image.Rect(0, 0, 240, 240))

	draw.Draw(board, board.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if (x%2 == 0 && y%2 == 0) || (x%2 != 0 && y%2 != 0) {
				for i := 0; i < 30; i++ {
					for j := 0; j < 30; j++ {
						board.Set(x*30+i, y*30+j, black)
					}
				}
			}
		}
	}

	file, err := os.Create("chessboard.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()

	png.Encode(file, board)
}
