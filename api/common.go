package api

import (
	"math/rand"
	"strings"
	"time"

	"github.com/Waifu-pics/waifu-api/config"
)

// CheckValid : check if endpoint is valid
func CheckValid(endpoint string, nsfw bool, conf config.Config) bool {
	var list = conf.Endpoints.Sfw
	if nsfw {
		list = conf.Endpoints.Nsfw
	}

	for _, v := range list {
		if v == endpoint {
			return true
		}
	}

	return false
}

// GetExtension : get file extension
func GetExtension(text string) string {
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

// RandomString : generate a random string
func RandomString(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_~"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// FindInSlice : find string in slice
func FindInSlice(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
