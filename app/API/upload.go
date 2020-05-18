package api

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
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

		resType := r.Header.Get("type")
		resToken := r.Header.Get("token")

		if !findInSlice(Config.ENDPOINTS, resType) {
			util.WriteResp(w, 400, "Invalid type!")
			return
		}

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

		md5sum := md5.Sum([]byte(buf.Bytes()))
		hash := hex.EncodeToString(md5sum[:])

		var mimetype = freq.Header.Get("Content-Type")
		var uplfilename = freq.Filename
		var filename = randomString(7) + "." + getExtension(uplfilename)

		count, _ := util.Database.Collection("uploads").CountDocuments(context.TODO(), bson.M{"md5": hash})
		if count > 0 {
			util.WriteResp(w, 400, "File already exists!")
			return
		}

		// Check if filename exists and reroll if yes
		filter := bson.M{"file": filename}
		count, _ = util.Database.Collection("uploads").CountDocuments(context.TODO(), filter)
		for count > 0 {
			filename = randomString(7) + "." + getExtension(uplfilename)
			count, _ = util.Database.Collection("uploads").CountDocuments(context.TODO(), filter)
		}

		// Actually upload file with S3
		err = util.Upload(buf, mimetype, filename, Config)
		if err != nil {
			util.WriteResp(w, 400, "File could not be uploaded!")
			return
		}

		// Check if user is an admin to skip verification
		if resToken != "" {
			count, _ = util.Database.Collection("admins").CountDocuments(context.TODO(), bson.M{"token": resToken})
			if count != 0 {
				util.Database.Collection("uploads").InsertOne(context.TODO(), bson.M{"file": filename, "verified": true, "md5": hash, "type": resType})
				return
			}
		}

		util.Database.Collection("uploads").InsertOne(context.TODO(), bson.M{"file": filename, "verified": false, "md5": hash, "type": resType})
		util.WriteResp(w, 200, "File uploaded!")

		defer r.Body.Close()
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

func findInSlice(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
