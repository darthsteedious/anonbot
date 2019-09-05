package clients

import (
	"encoding/json"
	"io"
)

type ApiResult struct {
	statusCode int
	body io.Reader
}

func (ar *ApiResult) Success() bool {
	return ar.statusCode >= 200 && ar.statusCode <= 400
}



type JsonApiResult struct {
	ApiResult
	data json.RawMessage
}
