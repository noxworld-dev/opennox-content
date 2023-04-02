package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

var (
	fIn   = flag.String("i", "tile.png", "input file name")
	fRows = flag.Int("rows", 8, "rows to split")
	fCols = flag.Int("cols", 12, "rows to split")
)

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	f, err := os.Open(*fIn)
	if err != nil {
		return err
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		return err
	}
	ext := filepath.Ext(*fIn)
	base := strings.TrimSuffix(*fIn, ext)
	dir := fmt.Sprintf("%s.%d", base, *fCols)
	if err = os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	sz := img.Bounds().Size()
	m := *fCols
	n := *fRows
	dx := sz.X / m
	dy := sz.Y / n
	part := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			draw.Draw(part, part.Rect, img, image.Pt(j*dx, i*dy), draw.Src)
			buf.Reset()
			if err = png.Encode(&buf, part); err != nil {
				return err
			}
			name := filepath.Join(dir, fmt.Sprintf("tile.%d.png", 100*i+j))
			if err = os.WriteFile(name, buf.Bytes(), 0644); err != nil {
				return err
			}
		}
	}
	return nil
}
