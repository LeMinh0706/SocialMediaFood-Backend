package middlewares

import (
	"strings"

	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeaderKey  = "authorization"
	AuthorizationPayloadKey = "authorization_payload"
)

// Update middle with less code
// Error is in Token package
// Use ErrNonKnow but knowing in advance :))
func AuthorizeMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			response.ErrorResponse(ctx, 40101)
			ctx.Abort()
			return
		}

		if !strings.HasPrefix(authorizationHeader, "Bearer") {
			authorizationHeader = "Bearer " + authorizationHeader
			ctx.Request.Header.Set(AuthorizationHeaderKey, authorizationHeader)
		}

		fields := strings.Fields(authorizationHeader)

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			response.ErrorResponse(ctx, 40101)
			ctx.Abort()
			return
		}
		ctx.Set(AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}
