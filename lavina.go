package lavina

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
