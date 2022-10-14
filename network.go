package lvn

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
func Res(statusCode int, data interface{}, message string) (int, string, []byte) {
	result := response{}
	result.Message = message
	result.Data = data
	result.IsOk = statusCode < 300 && statusCode >= 200
	bytes, _ := Marshal(result)

	return statusCode, "application/json", bytes
}
