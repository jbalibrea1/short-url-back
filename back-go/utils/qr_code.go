package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"

	qrcode "github.com/yeqown/go-qrcode"
)

// GenerateQRCode genera un código QR para una URL dada y lo devuelve como una cadena en base64.
func GenerateQRCode(url string) (string, error) {
	qrc, err := qrcode.New(url)
	if err != nil {
		return "", fmt.Errorf("could not create QR code: %v", err)
	}

	var buf bytes.Buffer
	// if err = qrc.Save("./a.png"); err != nil {
	// 	panic(err)
	// }
	// return "prueba", nil
	if err := qrc.SaveTo(io.Writer(&buf)); err != nil {
		return "", fmt.Errorf("could not encode QR code to PNG: %v", err)
	}

	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	return "data:image/png;base64," + encoded, nil
}
