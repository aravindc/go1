package main

import (
	"bytes"
	"image/png"
	"testing"
)

func TestGenerateQRCodeReturnsValue(t *testing.T) {
	buffer := new(bytes.Buffer)
	GenerateQRCode(buffer, "020 7152 4849")
	if buffer.Len() == 0 {
		t.Errorf("No QR Code generated")
	}
}

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
	buffer := new(bytes.Buffer)
	GenerateQRCode(buffer, "555-2368")
	_, err := png.Decode(buffer)

	if err != nil {
		t.Errorf("Generated QRCode is not a PNG: %s", err)
	}
}
