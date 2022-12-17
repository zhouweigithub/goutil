package qrcodutil

import (
	"image"
	"io"

	qrcode "github.com/skip2/go-qrcode"
)

// encodes, then writes a QR Code to the given filename in PNG format.
//
//	content:  content
//	size:     size is both the image width and height in pixels.
//	filename: saved filename(.png)
func CreatePngFile(content string, size int, filename string) error {
	return qrcode.WriteFile(content, qrcode.Medium, size, filename)
}

// returns the QR Code as an image.Image.
//
//	content:  content
//	size:     size is both the image width and height in pixels.
func CreateImage(content string, size int) (image.Image, error) {
	var q, err = qrcode.New(content, qrcode.Medium)
	if err != nil {
		return nil, err
	}
	return q.Image(size), nil
}

// returns the QR Code as a PNG image.
//
//	content:  content
//	size:     size is both the image width and height in pixels.
func CreatePngBytes(content string, size int) ([]byte, error) {
	var q, err = qrcode.New(content, qrcode.Medium)
	if err != nil {
		return nil, err
	}
	return q.PNG(size)
}

// writes the QR Code as a PNG image to io.Writer.
//
//	content:  content
//	size:     size is both the image width and height in pixels.
//	out:      io.Writer
func CreatePngIowriter(content string, size int, out io.Writer) error {
	var q, err = qrcode.New(content, qrcode.Medium)
	if err != nil {
		return err
	}
	return q.Write(size, out)
}
