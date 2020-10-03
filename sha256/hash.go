package sha256

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func Compute(data []byte) []byte {
	hash := sha256.Sum256(data)
	return ToHex(hash[:])
}

func ComputeHMAC(message, secret []byte) ([]byte, error) {
	h := hmac.New(sha256.New, secret)

	_, err := h.Write(message)
	if err != nil {
		return nil, err
	}

	return ToHex(h.Sum(nil)), nil
}

func ToHex(hash []byte) []byte {
	bs := make([]byte, hex.EncodedLen(len(hash)))
	hex.Encode(bs, hash)
	return bs
}
