package errc

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
)

func FromGrpcServiceError(err error) *Error {
	var errk *errors.Error
	if errors.As(err, &errk) {
		if value, exists := ErrCode_value[errk.Reason]; exists {
			return NewError(NewCode(int(value)), errk.Message)
		}
		return NewError(CodeUnknowErr, errk.Message)

	} else {
		errk := errors.FromError(err)
		if value, exists := ErrCode_value[errk.Reason]; exists {
			return NewError(NewCode(int(value)), errk.Message)
		}
		return NewError(CodeUnknowErr, errk.Message)
	}
}

func ToGrpcServiceError(err error, isDebug ...bool) *errors.Error {
	var errk *errors.Error
	if !errors.As(err, &errk) {
		var errc *Error
		if errors.As(err, &errc) {
			httpCode := 500
			switch errc.Code() {
			case CodeTokenExpired, CodeTokenInvalid, CodeTokenMissing, CodeUserNoAuth:
				httpCode = 403
			}

			if len(isDebug) > 0 && isDebug[0] {
				errk = errors.New(httpCode, errc.Code().String(), errc.Error())
			} else {
				errk = errors.New(httpCode, errc.Code().String(), errc.Message())
			}

		} else {
			errk = errors.New(500, "InnerError", err.Error())
		}
	}

	return errk
}

func ToHttpServiceError(err error, isDebug ...bool) *errors.Error {
	var errk *errors.Error
	if !errors.As(err, &errk) {
		var errc *Error
		if errors.As(err, &errc) {
			httpCode := 500
			switch errc.Code() {
			case CodeTokenExpired, CodeTokenInvalid, CodeTokenMissing, CodeUserNoAuth:
				httpCode = 403
			}

			if len(isDebug) > 0 && isDebug[0] {
				errk = errors.New(httpCode, fmt.Sprintf("%s;%d", errc.Code().String(), errc.Code().Number()), errc.Error())
			} else {
				errk = errors.New(httpCode, fmt.Sprintf("%s;%d", errc.Code().String(), errc.Code().Number()), errc.Message())
			}

		} else {
			errk = errors.New(500, fmt.Sprintf("%s;%d", CodeUnknowErr.String(), CodeUnknowErr.Number()), err.Error())
		}
	}

	return errk
}
