package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		duration := time.Since(start)
		logger.Info("Request",
			zap.String("method", ctx.Request.Method),
			zap.String("path", ctx.Request.URL.Path),
			zap.String("status", ctx.Request.URL.Path),
			zap.Int("status", ctx.Writer.Status()),
			zap.String("duration", duration.String()),
			zap.String("ip", ctx.ClientIP()),
		)
	}
}
