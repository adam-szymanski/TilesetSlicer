package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
)

func main() {
	// Declare flags.
	var filePath string
	flag.StringVar(&filePath, "file", "", "Path to image to be sliced.")
	var tileW int
	flag.IntVar(&tileW, "w", 0, "width of tile")
	var tileH int
	flag.IntVar(&tileH, "h", 0, "height of tile")
	var outDir string
	flag.StringVar(&outDir, "out", "out", "Output folder")
	var outName string
	flag.StringVar(&outName, "out_filename", "tile_", "Name of output file name.")
	flag.Parse()

	// Read tileset image.
	tiles, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("could not open file %s because %s", filePath, err.Error())
	}
	defer tiles.Close()
	src, _, err := image.Decode(tiles)
	if err != nil {
		panic(err.Error())
	}

	// Let's check necessary conditions before we start processing.
	bounds := src.Bounds()
	size := bounds.Size()
	if tileW == 0 || size.X%tileW != 0 {
		fmt.Printf("image width is %d, tile width of %d is not valid divider", size.X, tileW)
		os.Exit(1)
	}
	if tileH == 0 || size.Y%tileH != 0 {
		fmt.Printf("image height is %d, tile height of %d is not valid divider", size.Y, tileH)
		os.Exit(1)
	}
	if err := os.RemoveAll(outDir); err != nil {
		fmt.Printf("coud not clean output directory because: %s", err.Error())
		os.Exit(1)
	}

	// Creation of output dir.
	if err := os.Mkdir(outDir, os.ModeDir); err != nil {
		fmt.Printf("could not create output dir because: %s", err)
		os.Exit(1)
	}

	// Let's redraw tiles from tileset image into seperate image and save it into output directoru
	tileID := 0
	tilesW := size.X / tileW
	tilesH := size.Y / tileH

	offsetX := bounds.Min.X
	offsetY := bounds.Min.Y
	img := image.NewRGBA(image.Rect(0, 0, tileW, tileH))
	for yTile := 0; yTile < tilesH; yTile++ {
		for xTile := 0; xTile < tilesW; xTile++ {
			draw.Draw(img, image.Rect(0, 0, tileW, tileH), src, image.Point{offsetX, offsetY}, draw.Src)

			// Saving single tile
			outFile, err := os.Create(filepath.Join(outDir, fmt.Sprintf("%s%03d.png", outName, tileID)))
			if err != nil {
				fmt.Printf("could not create output file because: %s", err.Error())
				os.Exit(1)
			}
			defer outFile.Close()
			if err = png.Encode(outFile, img); err != nil {
				fmt.Printf("could not ecode output image because: %s", err.Error())
				os.Exit(1)
			}

			tileID++
			offsetX += tileW
		}
		offsetX = bounds.Min.X
		offsetY += tileH
	}
}
