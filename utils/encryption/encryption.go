package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

func NewEncryptionKey(secret string) *[32]byte {
	if len(secret) < 32 {
		secret = fmt.Sprintf("%-32v", secret)
	}

	if len(secret) > 32 {
		secret = secret[0:32]
	}

	buffer := []byte(secret)
	var key [32]byte

	for i := 0; i < 32; i++ {
		key[i] = buffer[i]
	}

	return &key
}

func Encrypt(plainText []byte, key *[32]byte) (cipherText []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plainText, nil), nil
}

func Decrypt(cipherText []byte, key *[32]byte) (plainText []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(cipherText) < gcm.NonceSize() {
		return nil, errors.New("malformed cipher text")
	}

	return gcm.Open(nil,
		cipherText[:gcm.NonceSize()],
		cipherText[gcm.NonceSize():],
		nil,
	)
}
