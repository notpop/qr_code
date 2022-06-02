package main

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

const TARGET_STRING = "https://twitter.com/getKishimoto"
const OUTPUT_FILE_NAME = "./result/qrcode.png"

const SCALE_WIDTH = 200
const SCALE_HEIGHT = 200

func main() {
	qrCode, _ := qr.Encode(TARGET_STRING, qr.M, qr.Auto)

	qrCode, _ = barcode.Scale(qrCode, SCALE_WIDTH, SCALE_HEIGHT)

	file, _ := os.Create(OUTPUT_FILE_NAME)
	defer file.Close()

	png.Encode(file, qrCode)
}
