package api

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"waifu.pics/util/file"
	"waifu.pics/util/web"
)

// UploadHandle : Handle uploads to the API
func (api API) UploadHandle(w http.ResponseWriter, r *http.Request) {
	const uploadSize int64 = 30 * 1000 * 1000

	resType := r.Header.Get("type")
	resToken := r.Header.Get("token")

	if !findInSlice(api.Config.ENDPOINTS, resType) {
		web.WriteResp(w, 400, "Invalid type!")
		return
	}

	fileForm, freq, err := r.FormFile("uploadFile")
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

	allowedTypes := []string{"image/jpeg", "image/png", "image/x-png"}

	// Check if file is actually an image
	if !findInSlice(allowedTypes, mimeType) || !findInSlice(allowedTypes, http.DetectContentType(buf.Bytes())) {
		web.WriteResp(w, 400, "File is not an image!")
		return
	}

	count, _ := api.Database.Collection("uploads").CountDocuments(context.TODO(), bson.M{"md5": hash})
	if count > 0 {
		web.WriteResp(w, 400, "File already exists!")
		return
	}

	// Check if filename exists and roll until its unique
	filter := bson.M{"file": filename}
	count, _ = api.Database.Collection("uploads").CountDocuments(context.TODO(), filter)
	for count > 0 {
		filename = randomString(7) + "." + getExtension(uplFileName)
		count, _ = api.Database.Collection("uploads").CountDocuments(context.TODO(), filter)
	}

	// Actually upload file with S3
	err = file.Upload(buf, mimeType, filename, api.Config)
	if err != nil {
		web.WriteResp(w, 400, "File could not be uploaded!")
		return
	}

	// Check if user is an admin to skip verification
	if resToken != "" {
		count, _ = api.Database.Collection("admins").CountDocuments(context.TODO(), bson.M{"token": resToken})
		if count != 0 {
			api.Database.Collection("uploads").InsertOne(context.TODO(), bson.M{"file": filename, "verified": true, "md5": hash, "type": resType})
			return
		}
	}

	_, err = api.Database.Collection("uploads").InsertOne(context.TODO(), bson.M{"file": filename, "verified": false, "md5": hash, "type": resType})
	if err != nil {
		_ = file.DeleteFile(filename, api.Config)
	}

	web.WriteResp(w, 200, "File uploaded!")

	defer r.Body.Close()
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
