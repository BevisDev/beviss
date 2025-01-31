package middleware

import (
	"github.com/gin-gonic/gin"
	"goauth/src/main/consts"
	"goauth/src/main/dto/response"
	"goauth/src/main/infrastructure/logger"
	"goauth/src/main/utils"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			state := utils.GetState(c)
			if err := recover(); err != nil {
				logger.Error(state, "Error occurred {}", err)
			}
			response.ServerError(c, consts.ServerError)
			c.Abort()
		}()

		c.Next()

		if len(c.Errors) != 0 {
			err := c.Errors.Last().Err
			response.SetErrMsg(c, http.StatusInternalServerError, consts.ServerError, err.Error())
		}
	}
}
