package sms

import (
	"encoding/json"
	"fmt"
	"strings"
)

type (
	Client struct {
		key    string
		secret string
	}
)

func New(key, secret string) *Client {
	return &Client{
		key:    key,
		secret: secret,
	}
}

func (c *Client) SendCode(template, phone string, codeLength int) (int, error) {
	phone = strings.Replace(phone, "+", "", -1)
	body := struct {
		Number string
		Length int
		Text   string
	}{
		Number: phone,
		Length: codeLength,
		Text:   template,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}

	bodyBytes, err = sendPost(bodyBytes, host+"/sms/sendCode", c.key, c.secret)
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

func (c *Client) VerifyCode(id int, code string) (bool, error) {
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

	bodyBytes, err = sendPost(bodyBytes, host+"/sms/verifyCode", c.key, c.secret)
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

func (c *Client) SendMessage(text, phone string) error {
	phone = strings.Replace(phone, "+", "", -1)
	body := struct {
		Number string
		Text   string
	}{
		Number: phone,
		Text:   text,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	bodyBytes, err = sendPost(bodyBytes, host+"/sms/send", c.key, c.secret)
	if err != nil {
		return err
	}

	resBody := struct {
		Data    string
		Message string
		IsOk    bool
	}{}

	err = json.Unmarshal(bodyBytes, &resBody)
	if err != nil {
		return err
	}

	if !resBody.IsOk {
		return fmt.Errorf("Recieved %v from server", resBody.Message)
	}

	return nil
}
