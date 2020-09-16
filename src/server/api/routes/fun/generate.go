package fun

import (
	"bufio"
	"bytes"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"

	"github.com/Riku32/waifu.pics/src/api"
	"github.com/jpoz/gomeme"
	"github.com/labstack/echo"
)

// GenerateMeme creates a meme using a waifu image.
// Why you ask? That is a very good question.
func (i route) GenerateMeme(c echo.Context) error {
	body := new(ReqGenerate)
	if err := c.Bind(body); err != nil {
		return c.JSON(400, api.Basic{Message: api.ErrInvalidJSON})
	}

	fileget, err := i.Database.GetFiles(body.Endpoint.Type, body.Endpoint.Nsfw, nil, 1)
	if err != nil || len(fileget) == 0 {
		return c.JSON(500, api.Basic{Message: api.ErrServer})
	}

	file, err := http.Get(i.Config.Web.Cdn + fileget[0])
	if err != nil {
		return c.JSON(500, api.Basic{Message: "Could not get image"})
	}
	defer file.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(file.Body)

	contentType := http.DetectContentType(buf.Bytes())

	// Create image config
	config := gomeme.NewConfig()
	config.TopText = body.Text.Top
	config.BottomText = body.Text.Bottom

	meme := &gomeme.Meme{
		Config: config,
	}

	switch contentType {
	case "image/gif":
		img, err := gif.DecodeAll(buf)
		if err != nil {
			return c.JSON(500, api.Basic{Message: ErrDecode})
		}
		config.FontSize = findWidth(img.Config.Width, len(biggerString(body.Text.Top, body.Text.Bottom)))
		meme.Memeable = gomeme.GIF{GIF: img}
	case "image/jpeg":
		img, err := jpeg.Decode(buf)
		if err != nil {
			return c.JSON(500, api.Basic{Message: ErrDecode})
		}
		config.FontSize = findWidth(img.Bounds().Dx(), len(biggerString(body.Text.Top, body.Text.Bottom)))
		meme.Memeable = gomeme.JPEG{Image: img}
	case "image/png":
		img, err := png.Decode(buf)
		if err != nil {
			return c.JSON(500, api.Basic{Message: ErrDecode})
		}
		config.FontSize = findWidth(img.Bounds().Dx(), len(biggerString(body.Text.Top, body.Text.Bottom)))
		meme.Memeable = gomeme.PNG{Image: img}
	default:
		return c.JSON(500, api.Basic{Message: "Invalid content type"})
	}

	var responseBuf bytes.Buffer
	writer := bufio.NewWriter(&responseBuf)

	err = meme.Write(writer)
	if err != nil {
		return c.JSON(500, api.Basic{Message: "Unable to write to buffer"})
	}

	return c.Blob(200, contentType, responseBuf.Bytes())
}

func findWidth(imgWidth int, textLength int) float64 {
	if textLength <= 2 {
		return float64(imgWidth) * 1 / float64(textLength)
	} else if textLength <= 4 {
		return float64(imgWidth) * 2 / float64(textLength)
	} else {
		return float64(imgWidth) * 3 / float64(textLength)
	}
}

func biggerString(A, B string) string {
	if len(A) > len(B) {
		return A
	}
	return B
}
