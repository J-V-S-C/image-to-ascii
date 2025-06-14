package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <image_path>")
		os.Exit(1)
	}
	imagePath := os.Args[1]

	fmt.Println("Choose a conversion mode:")
	fmt.Println("1 = ANSI")
	fmt.Println("2 = Colored ASCII")
	fmt.Println("3 = Plain ASCII")
	fmt.Print("Enter option (1, 2, or 3): ")

	var option string
	fmt.Scanln(&option)

	img := OpenImage(imagePath)
	for _, line := range ConvertImage(img, option) {
		fmt.Println(line)
	}
}

func OpenImage(fileName string) image.Image {
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("err opening file")
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Println("err decoding img")
	}

	return img
}

func ConvertImage(img image.Image, option string) []string {
	width := 150
	heigth := int(float64(img.Bounds().Dy()) / float64(img.Bounds().Dx()) * float64(width) * 0.55)
	resized := resize.Resize(uint(width), uint(heigth), img, resize.Lanczos3)
	var ansiImage []string

	for y := range resized.Bounds().Dy() {
		var row string
		for x := range resized.Bounds().Dx() {
			//replaceble func
			switch option {
			case "1":

				row += PixelToAnsi(resized.At(x, y))
				break
			case "2":
				row += PixelToColoredAscii(resized.At(x, y))
				break
			case "3":
				row += PixelToAscii(resized.At(x, y))
				break

			}
		}
		ansiImage = append(ansiImage, row)
	}
	return ansiImage
}

func PixelToAscii(c color.Color) string {
	const shapes = ".:-=+*#%@ "

	grayPixel := color.GrayModel.Convert(c).(color.Gray)
	pxl := int(grayPixel.Y) * (len(shapes) - 1) / 255

	return string(shapes[pxl])
}

func PixelToColoredAscii(c color.Color) string {
	const shadows = "@$B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/|()1{}[]?-_+~<>i!lI;:,\"^`'. "
	r, g, b, _ := c.RGBA()
	r, g, b = r>>8, g>>8, b>>8

	grayPixel := color.GrayModel.Convert(c).(color.Gray)
	pxl := int(grayPixel.Y) * (len(shadows) - 1) / 255

	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, string(shadows[len(shadows)-pxl-1]))
}

func PixelToAnsi(c color.Color) string {
	r, g, b, _ := c.RGBA()
	r, g, b = r>>8, g>>8, b>>8

	// color the background and the foreground
	return fmt.Sprintf("\033[38;2;%d;%d;%dm\033[48;2;%d;%d;%dm%s\033[0m",
		r, g, b, r, g, b, string('A'))
}
