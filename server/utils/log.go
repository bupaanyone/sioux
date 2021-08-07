package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

const contextKeyLogId = "ctx_log_id" // Key used to store log ID in a context.

func SetLogIdGin(c *gin.Context) {
	c.Set(contextKeyLogId, time.Now().UnixNano())
}

func GetLogId(ctx context.Context) int64 {
	if raw := ctx.Value(contextKeyLogId); raw != nil {
		if val, ok := raw.(int64); ok {
			return val
		}
	}
	return time.Now().UnixNano()
}

func Log(ctx context.Context, format string, args ...interface{}) {
	Output(2, ctx, format, args...)
}

func Output(callDepth int, ctx context.Context, format string, args ...interface{}) {
	args = append([]interface{}{GetLogId(ctx)}, args...)
	log.Output(callDepth+1, fmt.Sprintf("<%X> "+format, args...))
}
