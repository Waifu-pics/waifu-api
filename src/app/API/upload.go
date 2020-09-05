package api

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Riku32/waifu.pics/src/util/database"
	"github.com/Riku32/waifu.pics/src/util/file"
	"github.com/Riku32/waifu.pics/src/util/web"
	"github.com/gorilla/context"
)

const (
	uploadSize int64 = 30 * 1000 * 1000
)

var (
	allowedTypes = []string{"image/jpeg", "image/png", "image/x-png", "image/gif"}
)

func (api API) UploadHandle(w http.ResponseWriter, r *http.Request) {
	var res ImageRoute

	isAdmin := context.Get(r, "authbool").(bool)

	err := json.Unmarshal([]byte(r.FormValue("upload")), &res)
	if err != nil {
		web.SendJSON(w, 400, web.NewMessage(invalidjson))
		return
	}

	if !CheckValid(res, api.Config) {
		web.SendJSON(w, 400, web.NewMessage("invalid file type"))
		return
	}

	fileForm, freq, err := r.FormFile("upload")
	if err != nil {
		web.SendJSON(w, 400, web.NewMessage("file could not be uploaded"))
		return
	}

	if freq.Size >= uploadSize {
		web.SendJSON(w, 400, web.NewMessage("file too large"))
		return
	}

	defer fileForm.Close()

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(&io.LimitedReader{R: fileForm, N: uploadSize}); err != nil {
		web.SendJSON(w, 400, web.NewMessage("file could not be uploaded"))
		return
	}

	md5sum := md5.Sum(buf.Bytes())
	hash := hex.EncodeToString(md5sum[:])

	var mimeType = freq.Header.Get("Content-Type")
	var uplFileName = freq.Filename
	var filename = randomString(7) + "." + getExtension(uplFileName)

	// Check if file is actually an image
	if !findInSlice(allowedTypes, mimeType) || !findInSlice(allowedTypes, http.DetectContentType(buf.Bytes())) {
		web.SendJSON(w, 400, web.NewMessage("file is not an image"))
		return
	}

	err = api.Database.CreateFileInDB(filename, hash, res.Type, isAdmin, res.Nsfw)

	switch err {
	case nil:
		break
	case database.ErrorMD5Exists:
		web.SendJSON(w, 400, web.NewMessage("file already exists"))
		return
	case database.ErrorFileNameExists:
		var iter = 0
		for err != nil && iter < 10 {
			err = api.Database.CreateFileInDB(filename, hash, res.Type, false, res.Nsfw)
			if err != database.ErrorFileNameExists {
				web.SendJSON(w, 500, web.NewMessage(servererror))
				return
			}
			iter++
			if iter == 9 {
				web.SendJSON(w, 500, web.NewMessage(servererror))
				return
			}
		}
	default:
		web.SendJSON(w, 500, web.NewMessage(servererror))
		return
	}

	// Actually upload file with S3
	err = file.Upload(buf, mimeType, filename, api.Config)
	if err != nil {
		web.SendJSON(w, 400, web.NewMessage("file could not be uploaded"))
		return
	}

	web.SendJSON(w, 200, web.NewMessage("file uploaded"))

	defer r.Body.Close()
}
