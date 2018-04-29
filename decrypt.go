package gotpher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
)

// Decrypt gets bytes from an encrypted andOTP backup and it decrypts them.
func Decrypt(input []byte, password string) ([]byte, error) {
	key := sha256.Sum256([]byte(password))
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}
	c, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := input[:12]
	ciphertext := input[12:]
	return c.Open(nil, nonce, ciphertext, nil)
}
