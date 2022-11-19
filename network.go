package lvn

import "github.com/gin-gonic/gin"

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
	bytes, _ := marshal(result, omitKeys, selectKeys)

	return statusCode, "application/json", bytes
}
