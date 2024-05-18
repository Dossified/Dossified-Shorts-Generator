package image_conversion

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"net/http"
)

// ToJpeg converts a PNG image to JPEG format
func ToJpeg(imageBytes []byte) ([]byte, error) {

	// DetectContentType detects the content type
	contentType := http.DetectContentType(imageBytes)

	switch contentType {
	case "image/png":
		// Decode the PNG image bytes
		img, err := png.Decode(bytes.NewReader(imageBytes))

		if err != nil {
			return nil, err
		}

		buf := new(bytes.Buffer)

		// encode the image as a JPEG file
		if err := jpeg.Encode(buf, img, nil); err != nil {
			return nil, err
		}

		return buf.Bytes(), nil
	}

	return nil, fmt.Errorf("unable to convert %#v to jpeg", contentType)
}
