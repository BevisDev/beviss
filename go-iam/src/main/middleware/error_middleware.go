package middleware

import (
	"github.com/BevisDev/backend-template/utils"
	"github.com/gin-gonic/gin"
	"goiam/src/main/consts"
	"goiam/src/main/dto/response"
	"goiam/src/main/global"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			state := utils.GetState(c)
			if err := recover(); err != nil {
				global.Logger.Error(state, "Error occurred {}", err)
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
