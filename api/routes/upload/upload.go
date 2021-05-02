package upload

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Waifu-pics/waifu-api/api"
	"github.com/Waifu-pics/waifu-api/database"
	"github.com/labstack/echo"
)

const (
	uploadSize int64 = 30 * 1000 * 1000
)

var (
	allowedTypes = []string{"image/jpeg", "image/png", "image/x-png", "image/gif"}
)

// ErrFileNotUploaded : error when file cant be uploaded
var ErrFileNotUploaded = "file could not be uploaded"

// UploadHandle ; upload file endpoint
func (i Route) UploadHandle(c echo.Context) error {
	isAdmin := c.Get("authbool").(bool)

	var res api.ImageEndpoint

	err := json.Unmarshal([]byte(c.FormValue("upload")), &res)
	if err != nil {
		return c.JSON(400, api.Basic{Message: api.ErrInvalidJSON})
	}

	if !api.CheckValid(res.Type, res.Nsfw, i.Config) {
		return c.JSON(400, api.Basic{Message: "that is not a valid endpoint"})
	}

	form, err := c.FormFile("upload")
	if err != nil {
		return c.JSON(400, api.Basic{Message: ErrFileNotUploaded})
	}

	file, err := form.Open()
	if err != nil {
		return c.JSON(400, api.Basic{Message: ErrFileNotUploaded})
	}
	defer file.Close()

	if form.Size >= uploadSize {
		return c.JSON(400, api.Basic{Message: "file is too large"})
	}

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(&io.LimitedReader{R: file, N: uploadSize}); err != nil {
		return c.JSON(400, api.Basic{Message: ErrFileNotUploaded})
	}

	md5sum := md5.Sum(buf.Bytes())
	hash := hex.EncodeToString(md5sum[:])
	mimeType := form.Header.Get("Content-Type")
	filename := api.RandomString(7) + "." + api.GetExtension(form.Filename)

	// Check if file is actually an image
	if !api.FindInSlice(allowedTypes, mimeType) || !api.FindInSlice(allowedTypes, http.DetectContentType(buf.Bytes())) {
		return c.JSON(400, api.Basic{Message: "file is not an image"})
	}

	err = i.Database.CreateFileInDB(filename, hash, res.Type, isAdmin, res.Nsfw)
	switch err {
	case nil:
		break
	case database.ErrorMD5Exists:
		return c.JSON(400, api.Basic{Message: "file already exists"})
	case database.ErrorFileNameExists:
		var iter = 0
		for err != nil && iter < 10 {
			err = i.Database.CreateFileInDB(filename, hash, res.Type, isAdmin, res.Nsfw)
			if err != database.ErrorFileNameExists {
				return c.JSON(500, api.Basic{Message: api.ErrServer})
			}
			iter++
			if iter == 9 {
				return c.JSON(500, api.Basic{Message: api.ErrServer})
			}
		}
	default:
		fmt.Println("fa")
		return c.JSON(500, api.Basic{Message: api.ErrServer})
	}

	// Actually upload file with S3
	err = i.S3.UploadFile(*buf, filename, mimeType, true)
	if err != nil {
		fmt.Println(err)
		return c.JSON(400, api.Basic{Message: ErrFileNotUploaded})
	}

	return c.JSON(200, api.Basic{Message: fmt.Sprintf("%s has been uploaded", filename)})
}
