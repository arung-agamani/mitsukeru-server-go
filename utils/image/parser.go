package image

import (
	"bytes"
	"encoding/base64"
	"github.com/arung-agamani/mitsukeru-server-go/utils/logger"
	"image/png"
	"io"
	"strings"
)

func Base64ToPngBuffer(rawData string) (io.Reader, error) {
	trimmed := strings.TrimPrefix(rawData, "data:image/png;base64,")
	r := base64.NewDecoder(base64.StdEncoding, strings.NewReader(trimmed))
	img, err := png.Decode(r)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = png.Encode(buf, img)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	r = bytes.NewReader(buf.Bytes())
	return r, nil
}
