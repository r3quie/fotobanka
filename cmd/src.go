/*
Copyright © 2024 Pavel Kadera info@pavelkadera.cz

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"golang.org/x/term"

	"github.com/nfnt/resize"
)

// Prints given string line by line. Does not check whether argument is multiline
func PrintLicence(i string) {
	var lines []string
	if strings.Contains(i, "\r\n") {
		lines = strings.Split(i, "\r\n")
	} else {
		lines = strings.Split(i, "\n")
	}
	for _, l := range lines {
		fmt.Println(l)
		time.Sleep(100 * time.Millisecond)
	}
}

// Prints image with given name. Uses short ASCII gradient.
func PrintImg(i string) {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	f, err := os.Open("imgs/" + i)
	if err != nil {
		if runtime.GOOS == "windows" {
			f, err = os.Open(filepath.Join(os.Getenv("APPDATA"), "fotobanka", "imgs", i))
		} else {
			f, err = os.Open(filepath.Join(os.Getenv("HOME"), ".fotobanka", "imgs", i))
		}
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	defer f.Close()
	pixels := fileToPixels(f)
	for _, row := range pixels {
		for _, p := range row {
			fmt.Print(p)
		}
		fmt.Println()
	}
	fmt.Println("Pro zobrazení celého obrázku, držte klávesu CTRL a zmáčkněte klávesu +/- nebo použijte kolečko myši.")
}

// Helper functions

// fileToPixels converts an image file (io.Reader) to a 2D slice of strings representing the image
func fileToPixels(f io.Reader) [][]string {
	i, _, err := image.Decode(f)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fd := int(os.Stdout.Fd())
	width, _, err := term.GetSize(fd)
	if err != nil {
		panic(err)
	}

	bounds := i.Bounds()
	img := resize.Resize(
		uint(float64(width)*0.9),
		uint((float64(bounds.Dy())*(float64(width)*0.6)/float64(bounds.Dx()))*0.8),
		i, resize.Lanczos3,
	)

	b := img.Bounds()
	w, h := b.Dx(), b.Dy()
	var p [][]string
	for y := 0; y < h; y++ {
		var row []string
		for x := 0; x < w; x++ {
			row = append(row, luminanceToChar(PixelToLuminance(rgbaToPixel(img.At(x, y).RGBA()))))
		}
		p = append(p, row)
	}
	return p
}

// rgbaToPixel converts RGBA values to a Pixel struct
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{
		int(r / 257),
		int(g / 257),
		int(b / 257),
		int(a / 257),
	}
}

// PixelToLuminance converts a Pixel to a luminance value
func PixelToLuminance(p Pixel) int {
	return int(0.2126*float64(p.R) + 0.7152*float64(p.G) + 0.0722*float64(p.B))
}

// luminanceToChar converts a luminance value to a character
func luminanceToChar(l int) string {
	//const gradient = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'."
	const gradient = " .:-=+*#%@"
	return string(gradient[l*len(gradient)/256])
}

type Pixel struct {
	R int
	G int
	B int
	A int
}
