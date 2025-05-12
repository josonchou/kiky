package restful

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func errorResult(err error) *BodyResult[map[string]string] {
	bodyResult := &BodyResult[map[string]string]{
		Data:    map[string]string{},
		Code:    1,
		Message: err.Error(),
	}
	if bizErr, ok := err.(*HttpLogicError); ok {
		bodyResult.Code = bizErr.ErrorCode
		bodyResult.Message = bizErr.Message
	} else {
		bodyResult.Code = 1
	}

	return bodyResult
}

func okResult[T any](data T) *BodyResult[T] {
	return &BodyResult[T]{
		Data:    data,
		Code:    0,
		Message: "success",
	}
}

func OkJson[T any](w http.ResponseWriter, resp T, err error) {
	// var bodyResult *BodyResult[T]

	if err != nil {
		httpx.OkJson(w, errorResult(err))
	} else {
		httpx.OkJson(w, okResult[T](resp))
	}
}

func OnlyOKCtx(ctx context.Context, w http.ResponseWriter, err error) {
	OkJsonCtx(ctx, w, map[string]string{}, err)
}

func OnlyOK(w http.ResponseWriter, err error) {
	OkJson(w, map[string]string{}, err)
}

func OkJsonCtx[T any](ctx context.Context, w http.ResponseWriter, resp T, err error) {
	// var bodyResult *BodyResult[T]

	if err != nil {
		httpx.OkJsonCtx(ctx, w, errorResult(err))
	} else {
		httpx.OkJsonCtx(ctx, w, okResult[T](resp))
	}
}

func ErrorJson(w http.ResponseWriter, err error) {
	httpx.OkJson(w, errorResult(err))
}

func ErrorJsonCtx(ctx context.Context, w http.ResponseWriter, err error) {
	httpx.OkJsonCtx(ctx, w, errorResult(err))
}
