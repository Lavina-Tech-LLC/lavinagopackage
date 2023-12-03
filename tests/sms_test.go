package tests

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/Lavina-Tech-LLC/lavinagopackage/v2/conf"
	"github.com/Lavina-Tech-LLC/lavinagopackage/v2/sms"
	"github.com/joho/godotenv"
)

var (
	key        = ""
	secret     = ""
	phone      = ""
	message    = "Test sms from lavinagopackage. Your code is [code]"
	codeLength = 3
)

func init() {
	godotenv.Load(conf.GetPath() + "/.env")
	key = os.Getenv("KEY")
	secret = os.Getenv("SECRET")
	phone = os.Getenv("PHONE")
	msg := os.Getenv("MESSAGE")
	if msg != "" {
		message = msg
	}

}

func TestFull(t *testing.T) {
	id, err := sms.SendCode(message, phone, key, secret, codeLength)

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	fmt.Println("Please enter code from sms into sms.env")
	time.Sleep(15 * time.Second)
	godotenv.Load(conf.GetPath() + "/sms.env")
	code := os.Getenv("CODE")

	success, err := sms.VerifyCode(id, code, key, secret)
	fmt.Println(success)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestSend(t *testing.T) {
	id, err := sms.SendCode(message, phone, key, secret, codeLength)

	fmt.Println(id)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestVerify(t *testing.T) {
	id := 206504
	code := "805"

	success, err := sms.VerifyCode(id, code, key, secret)

	fmt.Println(success)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestPost(t *testing.T) {
	client := sms.New(key, secret)
	err := client.SendMessage("Test test test", phone)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
