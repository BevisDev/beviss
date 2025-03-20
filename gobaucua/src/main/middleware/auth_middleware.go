package middleware

import (
	"github.com/BevisDev/backend-template/helper"
	"github.com/gin-gonic/gin"
	"gobaucua/src/main/consts"
	"gobaucua/src/main/dto/response"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		if helper.IsNilOrEmpty(accessToken) {
			response.Unauthorized(c, consts.InvalidAccessToken)
			c.Abort()
			return
		}

		signature := c.GetHeader("signature")
		if helper.IsNilOrEmpty(signature) {
			response.Unauthorized(c, consts.InvalidSignature)
			c.Abort()
			return
		}

		c.Next()
	}
}
