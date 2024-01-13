package main

import (
	"fmt"
	"image"
	"math"
	"os"

	"image/draw"
	"image/png"
)

func main() {
	img1Path := "run.png"
	imgPaths := []string{img1Path, img1Path, img1Path, img1Path, img1Path, img1Path, img1Path, img1Path}

	imgs, err := openAllImages(imgPaths)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	col := int(math.Ceil(math.Pow(float64(len(imgs)), 0.5)))
	fmt.Printf("col: %d\n", col)

	startPoint := image.Point{0, 0}

	// 描画するキャンバス
	canvas := image.Rectangle{startPoint, image.Point{imgs[0].Bounds().Size().X * col, imgs[0].Bounds().Size().Y * col}}
	rgba := image.NewRGBA(canvas)

	for i, img := range imgs {
		x := -img.Bounds().Dx() * (i % col)
		y := -img.Bounds().Dy() * (i / col)
		drawPoint := image.Point{x, y}
		draw.Draw(rgba, canvas, img, drawPoint, draw.Src)
	}

	out, err := os.Create("out.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = png.Encode(out, rgba); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func openImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	return img, err
}

func openAllImages(paths []string) ([]image.Image, error) {
	imgs := make([]image.Image, len(paths))
	for i, path := range paths {
		img, err := openImage(path)
		if err != nil {
			return nil, err
		}
		imgs[i] = img
	}
	return imgs, nil
}
