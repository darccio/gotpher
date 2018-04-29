package gotpher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

// Encrypt gets bytes from an unencrypted andOTP backup and it encrypts them.
func Encrypt(input []byte, password string) ([]byte, error) {
	key := sha256.Sum256([]byte(password))
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}
	c, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, 12)
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}
	return c.Seal(nonce, nonce, input, nil), nil
}
