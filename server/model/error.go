package model

import (
	"context"
	"fmt"

	"github.com/bupaanyone/sioux/server/utils"
)

type Error struct {
	Code     int
	Message  string
	Detail   error
	HttpCode int
}

func (e *Error) Error() string {
	if e == nil || e.Code == 0 {
		return ""
	}
	return fmt.Sprintf("(%d) %+v", e.Code, e)
}

func E(ctx context.Context, err error, e *Error) error {
	if err == nil || e == nil {
		return err
	}

	utils.Output(2, ctx, "[ERROR] Error occurred, code=%d, message=%s, detail=%+v.", e.Code, e.Message, err)

	return &Error{
		Code:     e.Code,
		Message:  e.Message,
		Detail:   err,
		HttpCode: e.HttpCode,
	}
}
