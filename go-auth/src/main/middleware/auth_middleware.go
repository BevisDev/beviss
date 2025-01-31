package middleware

import (
	"github.com/BevisDev/backend-template/utils"
	"github.com/gin-gonic/gin"
	"goauth/src/main/consts"
	"goauth/src/main/dto/response"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		if utils.IsNilOrEmpty(accessToken) {
			response.Unauthorized(c, consts.InvalidAccessToken)
			c.Abort()
			return
		}

		signature := c.GetHeader("signature")
		if utils.IsNilOrEmpty(signature) {
			response.Unauthorized(c, consts.InvalidSignature)
			c.Abort()
			return
		}

		c.Next()
	}
}
