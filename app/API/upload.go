package API

import (
	"bytes"
	"crypto/md5"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"waifu.pics/util"
)

// UploadHandle : Handle uploads to the API
func UploadHandle(mux *mux.Router, Config util.Config) {
	mux.HandleFunc("/api/upload", func(w http.ResponseWriter, r *http.Request) {
		file, freq, err := r.FormFile("uploadFile")
		if err != nil {
			util.WriteResp(w, 400, "File could not be uploaded!")
			return
		}

		defer file.Close()

		var buf bytes.Buffer
		if _, err := buf.ReadFrom(&io.LimitedReader{R: file, N: 30 * 1000 * 1000}); err != nil {
			util.WriteResp(w, 400, "Could not upload file: "+err.Error())
			return
		}

		hash := md5.New()
		hash.Write([]byte(buf.Bytes()))

		var mimetype = freq.Header.Get("Content-Type")
		var uplfilename = freq.Filename

		var filename = randomString(7) + "." + getExtension(uplfilename)

		util.Upload(buf, mimetype, filename, Config)
	})
}

func getExtension(text string) string {
	pos := strings.Index(text, ".")
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(".")
	if adjustedPos >= len(text) {
		return ""
	}
	return text[adjustedPos:len(text)]
}

func randomString(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_~"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
