package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

func getSecret(secret []byte) []byte {
	base := make([]byte, 32)
	copy(base, secret)
	return base
}

func Encrypt(data, secret []byte) (string, error) {
	block, err := aes.NewCipher(getSecret(secret))
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(data))

	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], data)

	res := base64.StdEncoding.EncodeToString(cipherText)

	return res, nil
}

func Decrypt(encoded string, secret []byte) (plain []byte, err error) {
	cipherText, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return
	}

	//Create a new AES cipher with the key and encrypted message
	block, err := aes.NewCipher(getSecret(secret))

	//IF NewCipher failed, exit:
	if err != nil {
		return
	}

	//IF the length of the cipherText is less than 16 Bytes:
	if len(cipherText) < aes.BlockSize {
		err = errors.New("cipherText block size is too short!")
		return
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}
