package restful

type BodyResult[T any] struct {
	// Code is the Biz response code.
	Code int `json:"code"`
	// Message is the response message.
	Message string `json:"msg"`
	// Data is the response data.
	Data T `json:"data"`
}
