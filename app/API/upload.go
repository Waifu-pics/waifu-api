package api

import (
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"waifu.pics/util"
)

// UploadHandle : Handle uploads to the API
func UploadHandle(mux *mux.Router, Config util.Config) {
	mux.HandleFunc("/api/upload", func(w http.ResponseWriter, r *http.Request) {
		const uploadsize int64 = 30 * 1000 * 1000

		file, freq, err := r.FormFile("uploadFile")
		if err != nil {
			util.WriteResp(w, 400, "File could not be uploaded!")
			return
		}

		if freq.Size >= uploadsize {
			util.WriteResp(w, 400, "File too large!")
			return
		}

		defer file.Close()

		var buf bytes.Buffer
		if _, err := buf.ReadFrom(&io.LimitedReader{R: file, N: uploadsize}); err != nil {
			util.WriteResp(w, 400, "File could not be uploaded!"+err.Error())
			return
		}

		hash := md5.New()
		hash.Write([]byte(buf.Bytes()))

		var mimetype = freq.Header.Get("Content-Type")
		var uplfilename = freq.Filename
		var filename = randomString(7) + "." + getExtension(uplfilename)

		// Check if filename exists and reroll if yes
		filter := bson.M{"file": filename}
		count, _ := util.Database.Collection("uploads").CountDocuments(context.TODO(), filter)
		for count > 0 {
			filename = randomString(7) + "." + getExtension(uplfilename)
			count, _ = util.Database.Collection("uploads").CountDocuments(context.TODO(), filter)
			fmt.Println("rip")
		}

		// Actually upload file with S3
		err = util.Upload(buf, mimetype, filename, Config)
		if err != nil {
			util.WriteResp(w, 400, "File could not be uploaded!")
			return
		}

		util.Database.Collection("uploads").InsertOne(context.TODO(), bson.M{"file": filename, "verified": false})
		util.WriteResp(w, 200, "File uploaded!")
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
	return text[adjustedPos:]
}

func randomString(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_~"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
