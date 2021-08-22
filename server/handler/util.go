package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/bupaanyone/sioux/server/model"
)

type Time time.Time

var location *time.Location

func (t Time) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte("\"\""), nil
	}
	return []byte("\"" + time.Time(t).In(location).Format("2006-01-02 15:04:05") + "\""), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	res, err := time.ParseInLocation("2006-01-02 15:04:05", s, location)
	if err != nil {
		return err
	}
	*t = Time(res)
	return nil
}

func TimePtrToTimeTimePtr(t *Time) *time.Time {
	if t == nil {
		return nil
	}
	val := time.Time(*t)
	return &val
}

type requestPage struct {
	Skip  *int64 `json:"skip" binding:"omitempty,gte=0" example:"120"` // 跳过结果数
	Limit *int64 `json:"limit" binding:"omitempty,gte=0" example:"10"` // 本页大小
}

type responseError struct {
	Code    int    `json:"code" example:"1"`                       // 错误码，0为成功、其余为错误
	Message string `json:"message" example:"something went wrong"` // 错误信息
}

func returnError(c *gin.Context, err error) {
	if err == nil {
		c.JSON(http.StatusOK, responseOk)
		return
	}

	if e, ok := err.(*model.Error); ok {
		c.JSON(e.HttpCode, responseError{
			Code:    e.Code,
			Message: e.Message,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, responseError{
		Code:    1,
		Message: "服务器故障",
	})
}

type responseEmpty struct {
	Code int `json:"code" example:"0"` // 错误码，0为成功、其余为错误
}

var responseOk = responseEmpty{Code: 0}

func init() {
	location, _ = time.LoadLocation("Local")
}
