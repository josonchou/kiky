package restful

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HttpResponse[T any](w http.ResponseWriter, resp T, err error) {
	bodyResult := BodyResult[T]{}

	if err != nil {
		bodyResult.Message = err.Error()
		if bizErr, ok := err.(*HttpLogicError); ok {
			bodyResult.Code = bizErr.ErrorCode
			bodyResult.Message = bizErr.Message
		} else {
			bodyResult.Code = 1
		}
	} else {
		bodyResult.Code = 0
		bodyResult.Data = resp
		bodyResult.Message = "success"
	}
	httpx.OkJson(w, bodyResult)
}
