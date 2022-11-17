package tests

import (
	"testing"

	"github.com/Lavina-Tech-LLC/lavinagopackage/v2/crypto"
)

func TestEncryptDecrypt(t *testing.T) {

	secret := "MyCipherSecret"
	plainText := "Hello world"
	cipherText, err := crypto.Encrypt([]byte(plainText), []byte(secret))
	if err != nil {
		panic(err)
	}
	decipherText, err := crypto.Decrypt(cipherText, []byte(secret))
	if err != nil {
		panic(err)
	}
	res := []testsRes[string]{
		{
			Out:  string(decipherText),
			Want: plainText,
			Test: "encrypting and decrypting",
		},
	}
	check(res, t)
}
