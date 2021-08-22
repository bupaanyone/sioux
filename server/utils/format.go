package utils

import (
	"encoding/json"
	"time"
)

func StringToPtr(s string) *string {
	return &s
}

func TimeToPtr(t time.Time) *time.Time {
	return &t
}

// Parse an integer (interpreted as milliseconds) to a time.Duration.
func IntToTime(t int) time.Duration {
	return time.Duration(t) * time.Millisecond
}

// Parse an integer pointer (interpreted as milliseconds) to a time.Duration pointer.
func IntPtrToTimePtr(t *int) *time.Duration {
	if t == nil {
		return nil
	}
	res := IntToTime(*t)
	return &res
}

func DumpObject(obj interface{}) string {
	res, _ := json.MarshalIndent(obj, "", "  ")
	return string(res)
}
