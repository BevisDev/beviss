package middleware

import (
	"github.com/BevisDev/backend-template/helper"
	"github.com/gin-gonic/gin"
	"gobaucua/src/main/consts"
	"gobaucua/src/main/dto/response"
	"gobaucua/src/main/lib"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			state := helper.GetState(c)
			if err := recover(); err != nil {
				lib.Logger.Error(state, "Error occurred {}", err)
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
