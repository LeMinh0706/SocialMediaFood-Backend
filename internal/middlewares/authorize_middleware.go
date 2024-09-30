package middlewares

import (
	"errors"
	"fmt"
	"strings"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/token"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthorizeMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("Unauthorize")
			response.ErrorResponse(ctx, 401, err.Error())
			ctx.Abort()
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("Invalid Token")
			response.ErrorResponse(ctx, 401, err.Error())
			ctx.Abort()
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("Unsupported authorization type %s", authorizationType)
			response.ErrorResponse(ctx, 401, err.Error())
			ctx.Abort()
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			response.ErrorResponse(ctx, 401, err.Error())
			ctx.Abort()
			return
		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
