package sms

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	host = "https://tools.lavina.tech"
)

func SendCode(text, phoneNumber, key, secret string, codeLength int) (id int, err error) {
	phoneNumber = strings.Replace(phoneNumber, "+", "", -1)
	body := struct {
		Number string
		Length int
		Text   string
	}{
		Number: phoneNumber,
		Length: codeLength,
		Text:   text,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return
	}

	bodyBytes, err = sendPost(bodyBytes, host+"/sms/sendCode", key, secret)
	if err != nil {
		return 0, err
	}

	resBody := struct {
		Data    int
		Message string
		IsOk    bool
	}{}

	err = json.Unmarshal(bodyBytes, &resBody)
	if err != nil {
		return 0, err
	}

	if !resBody.IsOk {
		return 0, fmt.Errorf("Recieved %v from server", resBody.Message)
	}

	return resBody.Data, nil
}

func VerifyCode(id int, code, key, secret string) (bool, error) {
	body := struct {
		Id   int
		Code string
	}{
		Id:   id,
		Code: code,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return false, err
	}

	bodyBytes, err = sendPost(bodyBytes, host+"/sms/verifyCode", key, secret)
	if err != nil {
		return false, err
	}

	resBody := struct {
		Data    bool
		Message string
		IsOk    bool
	}{}

	err = json.Unmarshal(bodyBytes, &resBody)
	if err != nil {
		return false, err
	}

	if !resBody.IsOk {
		return false, fmt.Errorf("Recieved %v from server", resBody.Message)
	}

	return resBody.Data, nil
}

func sendPost(bodyBytes []byte, url, key, secret string) ([]byte, error) {

	req, err := http.NewRequest("POST", url, bytes.NewReader(bodyBytes))
	if err != nil {
		return []byte{}, err
	}

	signStr := string(bodyBytes) + secret
	generatedSum := md5.Sum([]byte(signStr))
	sign := hex.EncodeToString(generatedSum[:])

	req.Header.Add("k", key)
	req.Header.Add("s", sign)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return []byte{}, err
	}

	if res.StatusCode != 200 {
		return []byte{}, fmt.Errorf("Recieved %v from server", res.StatusCode)
	}

	return io.ReadAll(res.Body)
}
