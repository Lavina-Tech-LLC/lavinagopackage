package tests

import (
	"fmt"
	"testing"

	"github.com/Lavina-Tech-LLC/lavinagopackage/v2/sms"
)

const (
	key    = ""
	secret = ""
)

func TestSend(t *testing.T) {
	phone := "998913588881"
	message := "Test sms from lavinagopackage. Your code is [code]"
	codeLength := 6

	id, err := sms.SendCode(message, phone, key, secret, codeLength)

	fmt.Println(id)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestVerify(t *testing.T) {
	id := 165971
	code := "766376"

	success, err := sms.VerifyCode(id, code, key, secret)

	fmt.Println(success)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
