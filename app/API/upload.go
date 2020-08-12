package api

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/context"
	"waifu.pics/util/database"
	"waifu.pics/util/file"
	"waifu.pics/util/web"
)

const (
	uploadSize int64 = 30 * 1000 * 1000
)

var (
	allowedTypes = []string{"image/jpeg", "image/png", "image/x-png", "image/gif"}
)

func (api API) UploadHandle(w http.ResponseWriter, r *http.Request) {
	resType := r.Header.Get("type")
	isAdmin := context.Get(r, "authbool").(bool)

	if !findInSlice(api.Config.ENDPOINTS, resType) {
		web.WriteResp(w, 400, "Invalid type!")
		return
	}

	fileForm, freq, err := r.FormFile("upload")
	if err != nil {
		web.WriteResp(w, 400, "File could not be uploaded!")
		return
	}

	if freq.Size >= uploadSize {
		web.WriteResp(w, 400, "File too large!")
		return
	}

	defer fileForm.Close()

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(&io.LimitedReader{R: fileForm, N: uploadSize}); err != nil {
		web.WriteResp(w, 400, "File could not be uploaded!"+err.Error())
		return
	}

	md5sum := md5.Sum(buf.Bytes())
	hash := hex.EncodeToString(md5sum[:])

	var mimeType = freq.Header.Get("Content-Type")
	var uplFileName = freq.Filename
	var filename = randomString(7) + "." + getExtension(uplFileName)

	// Check if file is actually an image
	if !findInSlice(allowedTypes, mimeType) || !findInSlice(allowedTypes, http.DetectContentType(buf.Bytes())) {
		web.WriteResp(w, 400, "File is not an image!")
		return
	}

	err = api.Database.CreateFileInDB(filename, hash, resType, isAdmin)

	switch err {
	case nil:
		break
	case database.ErrorMD5Exists:
		web.WriteResp(w, 400, "File already exists!")
		return
	case database.ErrorFileNameExists:
		var iter = 0
		for err != nil && iter < 10 {
			err = api.Database.CreateFileInDB(filename, hash, resType, false)
			if err != database.ErrorFileNameExists {
				fmt.Println("lol")
				web.WriteResp(w, 500, "Error")
				return
			}
			iter++
			if iter == 9 {
				fmt.Println("lol")
				web.WriteResp(w, 500, "Error")
				return
			}
		}
	default:
		fmt.Println(err)
		web.WriteResp(w, 500, "Error")
		return
	}

	// Actually upload file with S3
	err = file.Upload(buf, mimeType, filename, api.Config)
	if err != nil {
		web.WriteResp(w, 400, "File could not be uploaded!")
		return
	}

	web.WriteResp(w, 200, "File uploaded!")

	defer r.Body.Close()
}
