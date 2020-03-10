package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

var (
	ErrInvalidData = errors.New("invalid data")
)

type Cipher struct {
	gcm cipher.AEAD
}

func NewCipher(key string) (*Cipher, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &Cipher{gcm}, nil
}

func (c *Cipher) Encrypt(data []byte) ([]byte, error) {
	nonce := make([]byte, c.gcm.NonceSize())

	_, err := io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return c.gcm.Seal(nonce, nonce, data, nil), nil
}

func (c *Cipher) Decrypt(data []byte) ([]byte, error) {
	nonceSize := c.gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, ErrInvalidData
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	return c.gcm.Open(nil, nonce, ciphertext, nil)
}
