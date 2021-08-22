package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

const ctxKeyLogId = "ctx_log_id" // Key used to store log ID in a context.

func CtxSetLogIdGin(c *gin.Context) {
	c.Set(ctxKeyLogId, time.Now().UnixNano())
}

func CtxGetLogId(ctx context.Context) int64 {
	if raw := ctx.Value(ctxKeyLogId); raw != nil {
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
	args = append([]interface{}{CtxGetLogId(ctx)}, args...)
	log.Output(callDepth+1, fmt.Sprintf("<%X> "+format, args...))
}
