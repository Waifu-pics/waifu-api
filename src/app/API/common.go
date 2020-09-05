package api

import (
	"math/rand"
	"strings"
	"time"

	"github.com/Riku32/waifu.pics/src/util/config"
)

const (
	invalidjson string = "invalid json provided"
	errorencode        = "unable to encode response"
	servererror        = "there was a server error when trying to process this request"
)

type ImageRoute struct {
	Type string `json:"type"`
	Nsfw bool   `json:"nsfw"`
}

func CheckValid(i ImageRoute, conf config.Config) bool {
	var list = conf.Endpoints.Sfw
	if i.Nsfw {
		list = conf.Endpoints.Nsfw
	}

	for _, v := range list {
		if v == i.Type {
			return true
		}
	}

	return false
}

func getExtension(text string) string {
	pos := strings.LastIndex(text, ".")
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
