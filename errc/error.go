package errc

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = NewError(CodeNotFound, "not found")

	ErrCarOffline          = NewError(CodeCarOffline, "car is offline")
	ErrCarNotSupportType   = NewError(CodeCarNotSupportType, "car not support type")
	ErrCarMeasureNotFound  = NewError(CodeCarMeasureNotFound, "car measure not found")
	ErrCarLowPower         = NewError(CodeCarLowPower, "car is low power")
	ErrCarNotInOperation   = NewError(CodeCarNotInOperation, "car is not operation state")
	ErrCarNotInMaintenance = NewError(CodeCarNotInMaintenance, "car is not maintenance state")
	ErrCarPoiUnreachable   = NewError(CodeCarPoiUnreachable, "car poi unreachable")
	ErrCarNotStopped       = NewError(CodeCarNotStopped, "car has not stopped")
	ErrCarUnlocked         = NewError(CodeCarUnlocked, "car unlocked")
	ErrCarEstopped         = NewError(CodeCarEstopped, "car is in emergency stop")
	ErrCarCoordUnknown     = NewError(CodeCarCoordUnknown, "car coordinates unknown")
	ErrCarNotSupportAuto   = NewError(CodeCarNotSupportAuto, "car not support auto")
	ErrCarNotInRoad        = NewError(CodeCarNotInRoad, "car not in road")
)

func NewCode(code int) Code {
	return Code(code)
}

func NewError(code Code, message ...string) *Error {
	if len(message) > 0 {
		return &Error{code: code, message: message[0]}
	}
	return &Error{code: code, message: ""}
}

func WithError(err error, code Code, message ...string) *Error {
	if len(message) > 0 {
		return &Error{code: code, message: message[0], inner: err}
	}
	return &Error{code: code, message: "", inner: err}
}

type Code int

func (c Code) Number() int {
	return int(c)
}

func (c Code) String() string {
	if name, exists := ErrCode_name[int32(c)]; exists {
		return name
	}

	return fmt.Sprintf("%d", int(c))
}

type Error struct {
	code    Code
	message string
	inner   error
}

func (e *Error) Code() Code {
	return e.code
}

func (e *Error) Message() string {
	return e.message
}

func (e *Error) Error() string {
	if e.inner != nil {
		return fmt.Sprintf("[%d]%s %s", e.code.Number(), e.message, e.inner.Error())
	}
	return fmt.Sprintf("[%d]%s", e.code.Number(), e.message)
}

func (e *Error) Unwrap() error {
	if e.inner != nil {
		return e.inner
	}
	return nil
}

func (e *Error) Is(err error) bool {
	var errk *Error
	if errors.As(err, &errk) {
		return e.Code().Number() == errk.Code().Number()
	}
	return false
}

func (e *Error) IsCode(code Code) bool {
	return e.Code().Number() == code.Number()
}
