package lvn

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	IsOk    bool        `json:"isOk"`
}

func Response(data interface{}, message string, status bool) response {
	result := response{}
	result.Message = message
	result.Data = data
	result.IsOk = status
	return result
}

// Response for using with *gin.Context.Data() body is forced as camelCase
func Res(statusCode int, data interface{}, message string, omitKeys ...string) (int, string, []byte) {
	return res(statusCode, data, message, omitKeys, []string{})
}

// Receives gin.Context and error, and if err!=nil aborts context with error
func GinErr(c *gin.Context, statusCode int, err error, message string) {
	if err != nil {
		c.Data(Res(statusCode, err.Error(), message))
		c.Abort()
		Logger.Error(err.Error())
		panic("lvn.GinErr panic")
	}
}

func ResSelected(statusCode int, data interface{}, message string, selectKeys ...string) (int, string, []byte) {
	return res(statusCode, data, message, []string{}, selectKeys)
}

// Response for using with *gin.Context.Data() body is forced as camelCase
func res(statusCode int, data interface{}, message string, omitKeys, selectKeys []string) (int, string, []byte) {
	result := response{}
	result.Message = message
	result.Data = data
	result.IsOk = statusCode < 300 && statusCode >= 200
	if len(selectKeys) != 0 {
		selectKeys = append(selectKeys, "isOk", "message", "data")
	}
	bytes, err := marshal(result, omitKeys, selectKeys)
	if err != nil {
		fmt.Printf("lvn.res marshall error: %s", err.Error())
	}

	return statusCode, "application/json", bytes
}

func GetSign(c *gin.Context, secret string) (string, error) {
	signStr := c.Request.Method + c.Request.URL.Path + Ternary(c.Request.URL.RawQuery, "?"+c.Request.URL.RawQuery, "")

	if c.Request.Method != "GET" {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			return "", err
		}
		signStr += string(body)
		c.Set("body", body)
	}

	signStr += secret

	sign := md5.Sum([]byte(signStr))

	return hex.EncodeToString(sign[:]), nil
}
