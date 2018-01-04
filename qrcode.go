package main

import (
	"image"
	"image/png"
	"io"
	"os"
)

func main() {

	file, _ := os.Create("qrcode.png")
	defer file.Close()

	GenerateQRCode(file, "020 7152 4849")
}

//GenerateQRCode does function to convert string to qr code
func GenerateQRCode(w io.Writer, code string) {
	img := image.NewRGBA(image.Rect(0, 0, 25, 25))
	_ = png.Encode(w, img)
}
