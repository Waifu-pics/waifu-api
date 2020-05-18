package crypto

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type passwordConfig struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

// GeneratePassword : Create a salted hash from a plaintext password
func GeneratePassword(password string) (string, error) {
	// Argon config
	var c = &passwordConfig{
		time:    1,
		memory:  64 * 1024,
		threads: 4,
		keyLen:  32,
	}

	// Generate a Salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.Key([]byte(password), salt, c.time, c.memory, c.threads, c.keyLen)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2i$v=%d$m=%d,t=%d,p=%d$%s$%s"
	full := fmt.Sprintf(format, argon2.Version, c.memory, c.time, c.threads, b64Salt, b64Hash)
	return full, nil
}

// ComparePassword : Check if a plaintext password matches the salted hash
func ComparePassword(password, hash string) (bool, error) {

	parts := strings.Split(hash, "$")

	c := &passwordConfig{}
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &c.memory, &c.time, &c.threads)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}
	c.keyLen = uint32(len(decodedHash))

	comparisonHash := argon2.Key([]byte(password), salt, c.time, c.memory, c.threads, c.keyLen)

	return (subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1), nil
}
